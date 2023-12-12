package env

import (
	"bufio"
	"os"
	"strings"
)

const (
	Name                     = "ttyd"
	Environment              = "/etc/PG.conf"
	PanabitPathKey           = "PGPATH"
	PanabitControlPanelIndex = "webmain"

	Ramdisk         = "/usr/ramdisk"
	ExtensionHome   = Ramdisk + "/app"
	ExtensionDir    = ExtensionHome + "/" + Name
	ExtensionLibDir = ExtensionDir + "/lib"

	ControlPanelHome      = Ramdisk + "/admin"
	ExtensionCgiDir       = ControlPanelHome + "/cgi-bin/App/" + Name
	ExtensionWebAssetsDir = ControlPanelHome + "/html/App/" + Name
)

var (
	StorageHome                  string
	ExtensionStorageDir          string
	ExtensionCgiStorageDir       string
	ExtensionWebAssetsStorageDir string
)

func Init() {
	StorageHome = "/usr/panabit"
	if v, err := find(PanabitPathKey); err != nil && v != "" {
		StorageHome = v
	}
	ExtensionStorageDir = StorageHome + "/app/" + Name
	ExtensionCgiStorageDir = ExtensionStorageDir + "/web/cgi"
	ExtensionWebAssetsStorageDir = ExtensionStorageDir + "/web/html"
}

func find(key string) (string, error) {
	f, err := os.Open(Environment)
	if err != nil {
		return "", err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if strings.HasPrefix(key, t) {
			if w := strings.Split(t, "="); len(w) == 2 {
				return w[1], nil
			} else {
				return "", nil
			}
		}
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return "", err
}
