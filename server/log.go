package server

import (
	"net/http"
	"io"
	"html/template"
)

//func saveLog(reader *bufio.Reader, id string) {
//	logMap[id] = make(chan string, 50)
//	for {
//		line, err := reader.ReadString('\n')
//		if err != nil {
//			break
//		}
//	RETRY:
//		select {
//		case logMap[id] <- line:
//		default:
//			<-logMap[id]
//			goto RETRY
//		}
//	}
//}
//
func showLog(v http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		t, err := template.ParseFiles("./view/log.html")
		if err != nil {
			io.WriteString(v, err.Error())
			return
		}
		id := r.Form.Get("id")
		data := map[string]interface{}{"id": id}
		t.Execute(v, data)
	}
}
