package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"proxy-web/server"
	"proxy-web/util"
)

var daemon = flag.Bool("d", true, "default run deamon")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *daemon {
		args := make([]string, 1)
		args[0] = "-d=false"
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}

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
			for _, v := range data {
				p, _ := os.FindProcess(v.ProcessId)
				p.Kill()
				p.Release()

			}
		}
	}()
	<-cleanupDone
}
