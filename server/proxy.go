package server

import (
	"proxy-web/utils"
	"net/http"
	"io"
	"html/template"
	"path"
	"time"
	"os"
	"strconv"
	"fmt"
	"strings"
	"github.com/snail007/goproxy/sdk/android-ios"
	"runtime"
	"path/filepath"
	"os/exec"
)

func add(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	command := r.Form.Get("command")
	autoStart := r.Form.Get("auto")
	keyFile := r.Form.Get("key_file")
	crtFile := r.Form.Get("crt_file")

	serviceId, err := utils.SaveParams(name, command, autoStart, keyFile, crtFile)
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}

	data := make(map[string]interface{})
	data["id"] = serviceId
	data["command"] = command
	data["auto_start"] = autoStart
	data["name"] = name
	data["status"] = "未开启"
	utils.ReturnJson("success", data, v)
}

func show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./view/index.html")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		autoStart := utils.NewConfig().GetAutoStart()
		data := map[string]interface{}{"auto_start": autoStart}
		t.Execute(w, data)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error
	r.ParseForm()
	id := r.Form.Get("id")

	if id == "0" {
		data, err = utils.GetAllParams()
	} else {
		data, err = utils.GetParamsById(id)
		if err != nil {
			utils.ReturnJson(err.Error(), "", w)
		}
	}
	utils.ReturnJson("success", data, w)
}

func link(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var command string
		var err error
		id := r.Form.Get("id")
		command, err = getCommand(id)
		if err != nil {
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		fmt.Println(command)
		errStr := proxy.Start(id, command)
		if errStr != "" {
			utils.ReturnJson(errStr, "", v)
			return
		}
		utils.ChangeParameterDataById(id, "已开启")
		utils.ReturnJson("success", "", v)
	}
}

func getCommand(id string) (command string, err error) {
	parameter, err := utils.GetParamsById(id)
	if err != nil {
		return "", err
	}

	command += parameter["command"].(string)
	command = strings.Replace(command, "  ", " ", -1)
	command = strings.Replace(command, "\n", "", -1)
	if parameter["key_file"].(string) != "" {
		command += " -K " + parameter["key_file"].(string)
	}
	if parameter["crt_file"].(string) != "" {
		command += " -C " + parameter["crt_file"].(string)
	}
	command += " --log " + parameter["log"].(string)
	s, err := os.Stat("./log/")
	if err != nil || !s.IsDir() {
		os.Mkdir("./log/", os.ModePerm)
	}
	return command, nil
}

func close(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	if id == "undefined" {
		utils.ReturnJson("id not found", "", v)
		return
	}
	err := utils.ChangeParameterDataById(id, "未开启")
	if err != nil {
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	proxy.Stop(id)
	utils.ReturnJson("success", "", v)
	return
}

func uploade(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		fileSuffix := path.Ext(head.Filename)
		if err != nil {
			utils.ReturnJson(err.Error(), "",  v)
			return
		}
		defer file.Close()
		t := time.Now().Unix()
		fw, err := os.Create("./upload/" + strconv.FormatInt(t, 10) + fileSuffix)
		defer fw.Close()
		if err != nil {
			utils.ReturnJson(err.Error(), "",  v)
			return
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		name := fw.Name()
		utils.ReturnJson("", name, v)
		return
	}
}

func update(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	name := r.Form.Get("name")
	command := r.Form.Get("command")
	autoStart := r.Form.Get("auto")
	keyFile := r.Form.Get("key_file")
	crtFile := r.Form.Get("crt_file")

	err := utils.UpdateParams(id, name, command, autoStart, keyFile, crtFile)
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	utils.ReturnJson("success", "", v)
}

func deleteParameter(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	err := utils.DeleteParam(id)
	if err != nil {
		utils.ReturnJson(err.Error(), "", v)
	}
	utils.ReturnJson("success", "", v)
}

func autoStart(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	auto := r.Form.Get("auto")
	var dir string

	switch runtime.GOOS {
	case "windows":
		if auto == "auto" {
			dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
			dir = strings.Replace(dir, "\\", "/", -1)
			command := `./config/autostart.exe enable -k proxy-web -n proxy-web -c`
			commandSlice := strings.Split(command, " ")
			commandSlice = append(commandSlice, dir+`/proxy-web.exe c:`)
			fmt.Println(commandSlice)
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			is_success := utils.NewConfig().UpdateAutoStart("true")
			if !is_success {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson("修改配置失败", "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		} else {
			command := `./config/autostart.exe disable -k proxy-web`
			commandSlice := strings.Split(command, " ")
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		}

	case "darwin":
		if auto == "auto" {
			command := `./config/autostart enable -k proxy -n proxy -c`
			commandSlice := strings.Split(command, " ")
			commandSlice = append(commandSlice, `echo \"autostart\">~/config/autostart.txt`)
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		} else {
			command := `./config/autostart disable -k "proxy"`
			commandSlice := strings.Split(command, " ")
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		}
	case "linux":
		if auto == "auto" {
			command := `./config/autostart enable -k proxy -n proxy -c`
			commandSlice := strings.Split(command, " ")
			commandSlice = append(commandSlice, `echo \"autostart\">~/config/autostart.txt`)
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		} else {
			command := `./config/autostart disable -k "proxy"`
			commandSlice := strings.Split(command, " ")
			cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(string(output), "", v)
				return
			}
			utils.ReturnJson("success", output, v)
			return
		}

	}
}
