package ttyd

import (
	"os/exec"
	"os/user"

	"xie.sh.cn/panabit-ttyd/v2/pkg/env"
)

const (
	ttyd = env.ExtensionBinaryDir + "/ttyd"
)

func Start() (*exec.Cmd, error) {
	cmd := exec.Command(ttyd, "-o", "-W", "bash")
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	cmd.Dir = u.HomeDir
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func Cleanup() {}
