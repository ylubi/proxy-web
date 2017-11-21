package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"proxy/server"
	"proxy/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initSet()
	server.StartServer()
	clean()
}

func initSet() {
	var n int
	db, err := sql.Open("sqlite3", "./db/sqlite/foo.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	res := db.QueryRow("select count(*) from sqlite_master where type = 'table' and name = 'parameter'", 1)
	res.Scan(&n)
	if n == 0 {
		//文件没找到表，创建
		sqlBytes, _ := ioutil.ReadFile("./db/proxy.sql")
		sqlString := string(sqlBytes)
		_, err := db.Exec(sqlString)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func clean() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for _ = range signalChan {
			data := util.GetParameterExistPid()
			for _, d := range data {
				p, _ := os.FindProcess(d.ProcessId)
				p.Kill()
				p.Release()
				util.PutParameterPidTo0(d.ProcessId)
			}
		}
	}()
	<-cleanupDone
}
