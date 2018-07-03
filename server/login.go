package server


//func login(v http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("./view/login.html")
//	if err != nil {
//		io.WriteString(v, err.Error())
//		return
//	}
//	t.Execute(v, nil)
//}
//
//func doLogin(v http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	username, password, err := util.NewConfig().GetUsernameAndPassword()
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	if (r.Form["username"][0] == username) && (r.Form["password"][0] == password) {
//		sess, _ := globalSessions.SessionStart(v, r)
//		defer sess.SessionRelease(v)
//		sessionId := sess.SessionID()
//		util.SaveSession(sessionId)
//		util.ReturnJson(200, "", "success", v)
//		return
//	}
//	util.ReturnJson(500, "", "login failed", v)
//}
