package server

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"proxy-web/util"
	"strings"
	"time"

	"github.com/astaxie/beego/session"
)

var logMap = make(map[string]chan string)
var globalSessions *session.Manager

func basicAuth(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, timeStamp, err := util.GetSession()
		if err != nil {
			login(w, r)
			return
		}
		t := time.Now()
		now := int(t.Unix())
		if timeStamp < now {
			login(w, r)
			return
		}
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)
		if sess.SessionID() == sessionId {
			http.HandlerFunc(handler).ServeHTTP(w, r)
			return
		} else {
			login(w, r)
			return
		}
	})
}

func StartServer() {
	AutoStart()
	time.Sleep(3 * time.Second)
	initSession()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/doLogin", doLogin)
	http.Handle("/", basicAuth(show))
	http.Handle("/add", basicAuth(add))
	http.Handle("/close", basicAuth(close))
	http.Handle("/link", basicAuth(link))
	http.Handle("/getData", basicAuth(getData))
	http.Handle("/showLog", basicAuth(showLog))
	http.Handle("/uploade", basicAuth(uploade))
	http.Handle("/delete", basicAuth(deleteParameter))
	http.Handle("/keygen", basicAuth(keygen))
	port, err := util.NewConfig().GetServerPort()
	if err != nil {
		log.Fatal("get port failure: ", err)
	}
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listen port failure", err)
	}
}

func AutoStart() {
	data := util.GetParameter()
	for _, value := range data {
		if value.Auto == "是" {
			command, err := getCommand(value.Id)
			if err != nil {
				log.Fatal(err.Error())
			}
			go autoRunCommand(command, value.Id)
		} else {
			util.ChangeParameterDataById(0, "未开启", value.Id)
		}
	}
}

func autoRunCommand(command string, id string) {
	cmdChan := make(chan int)
	commandList := strings.Split(command, " ")
	cmd := exec.Command(commandList[0], commandList[1:]...)
	//错误输出通道
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//异步等待是否返回错误
	reader := bufio.NewReader(stderr)
	go saveLog(reader, id)
	go waitProcess(cmd, cmdChan, id)
	second := time.After(3 * time.Second)

	//判断2秒内是否有channel返回，有则是失败，阻塞3秒以上则为成功
	select {
	case <-cmdChan:
		Output := getLog(id)
		fmt.Println(Output)
	case <-second:
		pid := cmd.Process.Pid
		err := util.ChangeParameterDataById(pid, "已开启", id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func initSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("file", sessionConfig)
	go globalSessions.GC()
	util.InitSession()
}
