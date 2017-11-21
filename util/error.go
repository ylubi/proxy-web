package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Code   int
	Output interface{}
	Pid    int
}

func ReturnJson(code, pid int, output string, v http.ResponseWriter) {
	r := Result{Code: code, Pid: pid, Output: output}
	res, err := json.Marshal(r)
	if err != nil {
		log.Fatal("json error")
	}
	io.WriteString(v, string(res))
}
