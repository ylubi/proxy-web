package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"proxy-web/procotol"
	"proxy-web/util"
	"strconv"
	"strings"
	"time"
)

var logMap = make(map[string]chan string)

func add(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, parameter, err := util.SaveParameter(r.Form)
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
	}
	util.ReturnJson(200, id, parameter, v)
}

func show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./view/index.html")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		t.Execute(w, nil)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error
	r.ParseForm()
	if r.Form["id"][0] == "0" {
		data = util.GetParameter()
	} else {
		data, err = util.GetParameterById(r.Form["id"][0])
		if err != nil {
			io.WriteString(w, err.Error())
		}
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, string(dataJson))
}

func link(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var command string
		var err error
		id := r.Form["id"][0]
		command, err = getCommand(id)
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		fmt.Println(command)
		runCommand(command, v, id)
	}
}

func getCommand(id string) (string, error) {
	parameter, err := util.GetParameterById(id)
	if err != nil {
		return "", err
	}
	var command string
	switch parameter.Protocol {
	case "http":
		command, err = procotol.GetHttpCommand(parameter)
	case "tcp":
		command, err = procotol.GetTcpCommand(parameter)
	case "socks":
		command, err = procotol.GetSocksCommand(parameter)
	case "udp":
		command, err = procotol.GetUdpCommand(parameter)
	case "server":
		command, err = procotol.GetServerCommand(parameter)
	case "client":
		command, err = procotol.GetClientCommand(parameter)
	case "bridge":
		command, err = procotol.GetBridgeCommand(parameter)
	default:
		err := fmt.Errorf("protocol parameter error")
		return "", err
	}
	fmt.Println(command)
	return command, nil
}

func runCommand(command string, v http.ResponseWriter, id string) {
	var Code int
	cmdChan := make(chan int)
	commandList := strings.Split(command, " ")
	cmd := exec.Command(commandList[0], commandList[1:]...)
	//错误输出通道
	stderr, err := cmd.StderrPipe()
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	err = cmd.Start()
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	//异步等待是否返回错误
	reader := bufio.NewReader(stderr)
	go saveLog(reader, id)
	go waitProcess(cmd, cmdChan, id)
	second := time.After(3 * time.Second)
	var stringPid string
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	//判断2秒内是否有channel返回，有则是失败，阻塞3秒以上则为成功
	select {
	case <-cmdChan:
		Code = 500
	case <-second:
		pid := cmd.Process.Pid
		err := util.ChangeParameterDataById(pid, "已开启", id)
		stringPid = strconv.Itoa(pid)
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		Code = 200
	}
	//进行输入流读取
	Output := getLog(id)
	util.ReturnJson(Code, stringPid, Output, v)
}

func saveLog(reader *bufio.Reader, id string) {
	logMap[id] = make(chan string, 10)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
	RETRY:
		select {
		case logMap[id] <- line:
		default:
			<-logMap[id]
			goto RETRY
		}
	}
}

func showLog(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["id"][0] == "undefined" || r.Form["id"][0] == "" {
			util.ReturnJson(500, "", "not found pid", v)
			return
		}
		res := getLog(r.Form["id"][0])
		if res == "" {
			time.Sleep(2 * time.Second)
		}
		util.ReturnJson(200, "", res, v)
	}
}

func close(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form["pid"][0] == "undefined" {
		util.ReturnJson(500, "", "pid not found", v)
		return
	}
	if r.Form["id"][0] == "undefined" {
		util.ReturnJson(500, "", "id not found", v)
		return
	}
	pid, err := strconv.Atoi(r.Form["pid"][0])
	p, err := os.FindProcess(pid)
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	err = util.ChangeParameterDataById(0, "未开启", r.Form["id"][0])
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	err = p.Kill()
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	err = p.Release()
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	delete(logMap, r.Form["id"][0])
	util.ReturnJson(200, "", "success", v)
	return
}

func uploade(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		fileSuffix := path.Ext(head.Filename)
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		defer file.Close()
		t := time.Now().Unix()
		fw, err := os.Create("./static/upload/" + strconv.FormatInt(t, 10) + fileSuffix)
		defer fw.Close()
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		name := fw.Name()
		util.ReturnJson(200, "", name, v)
		return
	}
}

func getLog(id string) string {
	var log string
	output := ""
	for i := 0; i <= 10; i++ {
		select {
		case log = <-logMap[id]:
			output += log
		case <-time.After(1 * time.Second):
			return output
		}
	}
	return output
}

func waitProcess(cmd *exec.Cmd, cmdChan chan int, id string) {
	cmd.Wait()
	cmdChan <- 1
	time.Sleep(1 * time.Second)
	delete(logMap, id)
}

func deleteParameter(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form["id"][0]
	err := util.DeleteParameterDataById(id)
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
	}
	delete(logMap, id)
	util.ReturnJson(200, "", "success", v)
}

func StartServer() {
	AutoStart()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/", show)
	http.HandleFunc("/add", add)
	http.HandleFunc("/close", close)
	http.HandleFunc("/link", link)
	http.HandleFunc("/getData", getData)
	http.HandleFunc("/showLog", showLog)
	http.HandleFunc("/uploade", uploade)
	http.HandleFunc("/delete", deleteParameter)
	port, err := util.GetServerPort()
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
	case <-second:
		pid := cmd.Process.Pid
		err := util.ChangeParameterDataById(pid, "已开启", id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
