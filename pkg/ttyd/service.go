package ttyd

import (
	"os/exec"

	"xie.sh.cn/panabit-ttyd/v2/pkg/env"
)

const (
	ttyd = env.ExtensionBinaryDir + "/ttyd"
)

func Run() *exec.Cmd {
	cmd := exec.Command(ttyd, "-o")
	if err := cmd.Start(); err != nil {
		return nil
	}
	return cmd
}

func Cleanup() {}
