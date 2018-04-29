package server

import (
	"time"
	"bufio"
	"net/http"
	"proxy-web/util"
)

func getLog(id string) string {
	var log string
	output := ""
	for i := 0; i <= 20; i++ {
		select {
		case log = <-logMap[id]:
			output += log
		case <-time.After(1 * time.Second):
			return output
		}
	}
	return output
}

func saveLog(reader *bufio.Reader, id string) {
	logMap[id] = make(chan string, 50)
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
		id := r.Form.Get("id")
		if id == "undefined" || id == "" {
			util.ReturnJson(500, "", "not found pid", v)
			return
		}
		res := getLog(id)
		if res == "" {
			time.Sleep(2 * time.Second)
		}
		util.ReturnJson(200, "", res, v)
	}
}
