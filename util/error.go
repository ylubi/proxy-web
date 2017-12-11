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
	Id     string
}

func ReturnJson(code int, id, output string, v http.ResponseWriter) {
	r := Result{Code: code, Id: id, Output: output}
	res, err := json.Marshal(r)
	if err != nil {
		log.Fatal("json error")
	}
	io.WriteString(v, string(res))
}
