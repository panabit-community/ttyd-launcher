package env

const (
	Name             = "ttyd"
	Ramdisk          = "/usr/ramdisk"
	ExtensionHome    = Ramdisk + "/app"
	ControlPanelHome = Ramdisk + "/admin"
	ExtensionDir     = ExtensionHome + "/" + Name

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
	// TODO assign from /etc/PG.conf
	ExtensionStorageDir = StorageHome + "/app/" + Name
	ExtensionCgiStorageDir = ExtensionStorageDir + "/web/cgi"
	ExtensionWebAssetsStorageDir = ExtensionStorageDir + "/web/html"
}
