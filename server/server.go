package server

import (
	"log"
	"net/http"
	"github.com/astaxie/beego/session"
	"proxy-web/utils"
	"fmt"
	"github.com/snail007/goproxy/sdk/android-ios"
	"strings"
	"os"
)

//var logMap = make(map[string]chan string)
var globalSessions *session.Manager

func basicAuth(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
}

func StartServer() {
	AutoStart()
	InitShowLog()
	initSession()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", basicAuth(show))
	http.Handle("/add", basicAuth(add))
	http.Handle("/update", basicAuth(update))
	http.Handle("/close", basicAuth(close))
	http.Handle("/link", basicAuth(link))
	http.Handle("/getData", basicAuth(getData))
	http.Handle("/uploade", basicAuth(uploade))
	http.Handle("/delete", basicAuth(deleteParameter))
	http.Handle("/auto_start", basicAuth(autoStart))
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/doLogin", doLogin)
	//http.Handle("/keygen", basicAuth(keygen))
	port, err := utils.NewConfig().GetServerPort()
	fmt.Println(port)
	if err != nil {
		log.Fatal("get port failure: ", err)
	}
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listen port failure", err)
	}
}

func AutoStart() {
	datas, err := utils.InitParams()
	if err != nil {
		return
	}
	for _, data := range datas {
		var command string
		command += data["command"].(string)
		command = strings.Replace(command, "  ", " ", -1)
		command = strings.Replace(command, "\n", "", -1)
		if data["key_file"].(string) != "" {
			command += " -K " + data["key_file"].(string)
		}
		if data["crt_file"].(string) != "" {
			command += " -C " + data["crt_file"].(string)
		}
		command += " --log " + data["log"].(string)
		s, err := os.Stat("./log/")
		if err != nil || !s.IsDir() {
			os.Mkdir("./log/", os.ModePerm)
		}
		go autoRunCommand(data["id"].(string), command)
	}
}

func autoRunCommand(id, command string) {
	fmt.Println(command)
	errStr := proxy.Start(id, command)
	if errStr != "" {
		utils.ChangeParameterDataById(id, "未开启")
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
}
