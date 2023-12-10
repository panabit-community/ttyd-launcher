package start

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ttyd/v2/pkg/env"

	"github.com/google/subcommands"
)

type Cmd struct{}

func (*Cmd) Name() string     { return "start" }
func (*Cmd) Synopsis() string { return "Start the extension." }
func (*Cmd) Usage() string    { return "start" }

func (p *Cmd) SetFlags(_ *flag.FlagSet) {}

func (p *Cmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	env.CopyDir(env.ExtensionCgiStorageDir, env.ExtensionCgiDir, 644)
	env.CopyDir(env.ExtensionWebAssetsStorageDir, env.ExtensionWebAssetsDir, 644)
	return subcommands.ExitSuccess
}
