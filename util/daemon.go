package util

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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
