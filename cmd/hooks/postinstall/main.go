package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"xie.sh.cn/panabit-ttyd/v2/pkg/env"
)

func main() {
	p, err := filepath.Abs("./appctrl")
	if err != nil {
		os.Exit(1)
	}
	cmd := exec.Command(p, "start")
	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}
}

func init() {
	env.Init()
}
