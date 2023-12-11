package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"xie.sh.cn/panabit-ttyd/v2/pkg/ttyd"
)

var (
	mutex  sync.Mutex
	status bool
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				if !status {
					Run()
				}
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Printf("Received signal: %v, exiting.", s)
				os.Exit(0)
			}
		}
	}
}

func Run() {
	mutex.Lock()
	status = true
	defer func() {
		mutex.Unlock()
		status = false
	}()
	ttyd.Run()
}
