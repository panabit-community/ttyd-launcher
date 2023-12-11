package ttyd

import (
	"os/exec"

	"xie.sh.cn/panabit-ttyd/v2/pkg/env"
)

const (
	ttyd = env.ExtensionLibDir + "/ttyd"
)

func Run() {
	cmd := exec.Command(ttyd, "-o")
	if err := cmd.Start(); err != nil {
		return
	}
	cmd.Wait()
}

func Cleanup() {}
