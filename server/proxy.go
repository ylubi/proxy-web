package server

import (
	"proxy-web/util"
	"time"
	"os"
	"strconv"
	"io"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"runtime"
	"path"
	"html/template"
)

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
	id := r.Form.Get("id")
	if id == "0" {
		data = util.GetParameter()
	} else {
		data, err = util.GetParameterById(id)
		if err != nil {
			util.ReturnJson(501, "", err.Error(), w)
		}
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		util.ReturnJson(501, "", err.Error(), w)
		return
	}
	util.ReturnJson(200, "", string(dataJson), w)
}

func link(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var command string
		var err error
		id := r.Form.Get("id")
		command, err = getCommand(id)
		if err != nil {
			util.ReturnJson(500, "", err.Error(), v)
			return
		}
		fmt.Println(command)
		runCommand(command, v, id)
	}
}

func getCommand(id string) (command string, err error) {
	parameter, err := util.GetParameterById(id)
	if err != nil {
		return "", err
	}

	config := util.NewConfig()
	serverPath, err := config.GetServerPath()
	if err != nil {
		return
	}
	command += serverPath + "proxy "
	command += parameter.Params
	command = strings.Replace(command, "  ", " ", -1)
	command = strings.Replace(command, "\n", "", -1)
	if parameter.Key != "" {
		command += " -K " + parameter.Key
	}
	if parameter.Key != "" {
		command += " -C " + parameter.Crt
	}
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

func close(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pid := r.Form.Get("pid")
	id := r.Form.Get("id")
	if pid == "undefined" {
		util.ReturnJson(500, "", "pid not found", v)
		return
	}
	if id == "undefined" {
		util.ReturnJson(500, "", "id not found", v)
		return
	}
	err := util.ChangeParameterDataById(0, "未开启", id)
	delete(logMap, id)
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
	pidInt, err := strconv.Atoi(pid)
	p, err := os.FindProcess(pidInt)
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
	util.ReturnJson(200, "", "success", v)
	return
}

func keygen(v http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	if os != "linux" {
		util.ReturnJson(500, "", "os error", v)
		return
	}
	fmt.Println(os)
	path, err := util.NewConfig().GetServerPath()
	command := path + "proxy keygen"
	commandList := strings.Split(command, " ")
	cmd := exec.Command(commandList[0], commandList[1:]...)
	err = cmd.Run()

	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
		return
	}
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
		fw, err := os.Create("./upload/" + strconv.FormatInt(t, 10) + fileSuffix)
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

func deleteParameter(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	err := util.DeleteParameterDataById(id)
	if err != nil {
		util.ReturnJson(500, "", err.Error(), v)
	}
	delete(logMap, id)
	util.ReturnJson(200, "", "success", v)
}

func waitProcess(cmd *exec.Cmd, cmdChan chan int, id string) {
	cmd.Wait()
	cmdChan <- 1
	time.Sleep(1 * time.Second)
	delete(logMap, id)
}
