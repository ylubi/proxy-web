package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"proxy/procotol"
	"proxy/util"
	"strconv"
	"strings"
	"time"
)

var logMap = make(map[int]chan string)

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
	data := util.GetParameterExistPid()
	for index, d := range data {
		//有bug进程已经终止，还是能找到该进程
		_, err := os.FindProcess(d.ProcessId)
		if err != nil {
			delete(data, index)
			util.PutParameterPidTo0(d.ProcessId)
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
		var command string
		var err interface{}
		r.ParseForm()
		switch r.Form["protocol"][0] {
		case "http":
			command, err = procotol.GetHttpCommand(r.Form)
		case "tcp":
			command, err = procotol.GetTcpCommand(r.Form)
		case "socks":
			command, err = procotol.GetSocksCommand(r.Form)
		case "udp":
			command, err = procotol.GetUdpCommand(r.Form)
		case "tserver":
			command, err = procotol.GetTserverCommand(r.Form)
		case "tclient":
			command, err = procotol.GetTclientCommand(r.Form)
		case "tbridge":
			command, err = procotol.GetTbridgeCommand(r.Form)
		default:
			res := &util.Result{Code: 500, Output: "protocol parameter error"}
			resJson, _ := res.ReturnJson()
			io.WriteString(v, resJson)
			return
		}
		if err != nil {
			res := &util.Result{Code: 500, Output: "protocol error"}
			resJson, _ := res.ReturnJson()
			io.WriteString(v, resJson)
			return
		}
		fmt.Println(command)
		runCommand(command, v, r.Form)
	}
}

func runCommand(command string, v http.ResponseWriter, data url.Values) {
	cmdChan := make(chan int)
	commandList := strings.Split(command, " ")
	cmd := exec.Command(commandList[0], commandList[1:]...)
	res := &util.Result{}
	//错误输出通道
	stderr, err := cmd.StderrPipe()
	if err != nil {
		res = &util.Result{Code: 500, Output: err.Error()}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	err = cmd.Start()
	//异步等待是否返回错误
	reader := bufio.NewReader(stderr)
	pid := cmd.Process.Pid
	go saveLog(reader, pid)
	go waitProcess(cmd, cmdChan, pid)
	second := time.After(2 * time.Second)

	//判断2秒内是否有channel返回，有则是失败，阻塞1秒以上则为成功
	select {
	case <-cmdChan:
		res.Code = 500
	case <-second:
		util.SaveParameterByPid(data, pid)
		res.Code = 200
		res.Pid = pid
	}
	//进行输入流读取
	res.Output = getLog(pid)
	resJson, err := res.ReturnJson()
	if err != nil {
		res = &util.Result{Code: 500, Output: err.Error()}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	io.WriteString(v, string(resJson))
}

func saveLog(reader *bufio.Reader, pid int) {
	logMap[pid] = make(chan string, 10)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
	RETRY:
		select {
		case logMap[pid] <- line:
		default:
			<-logMap[pid]
			goto RETRY
		}
	}
}

func showLog(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["pid"][0] == "undefined" {
			io.WriteString(v, "not found pid")
			return
		}
		pid, err := strconv.Atoi(r.Form["pid"][0])
		if err != nil {
			io.WriteString(v, err.Error())
			return
		}
		res := getLog(pid)
		result := &util.Result{Code: 200, Output: res}
		resJson, _ := result.ReturnJson()
		io.WriteString(v, resJson)
	}
}

func close(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form["pid"][0] == "undefined" {
		res := &util.Result{Code: 500, Output: "pid not found"}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	pid, err := strconv.Atoi(r.Form["pid"][0])
	p, err := os.FindProcess(pid)
	if err != nil {
		res := &util.Result{Code: 500, Output: err.Error()}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	util.PutParameterPidTo0(pid)
	err = p.Kill()
	if err != nil {
		res := &util.Result{Code: 500, Output: err.Error()}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	err = p.Release()
	if err != nil {
		res := &util.Result{Code: 500, Output: err.Error()}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
	delete(logMap, pid)
	res := &util.Result{Code: 200, Output: "success"}
	resJson, _ := res.ReturnJson()
	io.WriteString(v, resJson)
	return
}

func uploade(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		fileSuffix := path.Ext(head.Filename)
		if err != nil {
			res := &util.Result{Code: 500, Output: err.Error()}
			resJson, _ := res.ReturnJson()
			io.WriteString(v, resJson)
			return
		}
		defer file.Close()
		t := time.Now().Unix()
		fw, err := os.Create("./static/upload/" + strconv.FormatInt(t, 10) + fileSuffix)
		defer fw.Close()
		if err != nil {
			res := &util.Result{Code: 500, Output: err.Error()}
			resJson, _ := res.ReturnJson()
			io.WriteString(v, resJson)
			return
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			res := &util.Result{Code: 500, Output: err.Error()}
			resJson, _ := res.ReturnJson()
			io.WriteString(v, resJson)
			return
		}
		name := fw.Name()
		res := &util.Result{Code: 200, Output: name}
		resJson, _ := res.ReturnJson()
		io.WriteString(v, resJson)
		return
	}
}

func getLog(pid int) string {
	var log string
	output := ""
	for i := 0; i <= 10; i++ {
		select {
		case log = <-logMap[pid]:
			output += log
		case <-time.After(1 * time.Second):
			return output
		}
	}
	return output
}

func waitProcess(cmd *exec.Cmd, cmdChan chan int, pid int) {
	cmd.Wait()
	cmdChan <- 1
	time.Sleep(1 * time.Second)
	delete(logMap, pid)
}

func StartServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", show)
	http.HandleFunc("/close", close)
	http.HandleFunc("/link", link)
	http.HandleFunc("/getData", getData)
	http.HandleFunc("/showLog", showLog)
	http.HandleFunc("/uploade", uploade)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen port failure: ", err)
	}
}
