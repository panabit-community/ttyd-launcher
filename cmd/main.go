package main

import (
	"xie.sh.cn/panabit-ttyd/v2/pkg/cmd/start"
	"xie.sh.cn/panabit-ttyd/v2/pkg/env"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&start.Cmd{}, "")
}

func init() {
	env.Init()
}
