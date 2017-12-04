package main

import (
	"os"
	"os/signal"
	"syscall"

	"proxyWebApplication/server"
	"proxyWebApplication/util"
)

func main() {
	server.StartServer()
	clean()
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
			data := util.GetParameter()
			for _, d := range data {
				p, _ := os.FindProcess(d.ProcessId)
				p.Kill()
				p.Release()
				util.DeleteParameterByPid(d.ProcessId)
			}
		}
	}()
	<-cleanupDone
}
