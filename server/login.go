package server

import (
	"net/http"
	"io"
	"proxy-web/utils"
	"html/template"
	"proxy/util"
)

func login(v http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./view/login.html")
	if err != nil {
		io.WriteString(v, err.Error())
		return
	}
	t.Execute(v, nil)
}

func doLogin(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username, password, err := util.NewConfig().GetUsernameAndPassword()
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	sess, _ := globalSessions.SessionStart(v, r)
	newSessionId := sess.SessionID()
	if lock && sessionId != newSessionId {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson("已有人登陆", "", v)
		return
	}
	if (r.Form["username"][0] == username) && (r.Form["password"][0] == password) {
		sessionId = sess.SessionID()
		defer sess.SessionRelease(v)
		utils.ReturnJson("success", "", v)
		return
	}

	v.WriteHeader(http.StatusInternalServerError)
	utils.ReturnJson("登陆失败", "", v)
}
