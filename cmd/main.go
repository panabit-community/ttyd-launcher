package main

import (
	"context"
	"flag"
	"os"

	"xie.sh.cn/panabit-ttyd/v2/pkg/cmd/start"
	"xie.sh.cn/panabit-ttyd/v2/pkg/cmd/status"
	"xie.sh.cn/panabit-ttyd/v2/pkg/env"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&start.Cmd{}, "")
	subcommands.Register(&status.Cmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

func init() {
	env.Init()
}
