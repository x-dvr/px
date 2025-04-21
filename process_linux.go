package px

import (
	"syscall"
)

func setNewProcessGroupAttr(attr *syscall.SysProcAttr) {
	attr.Setpgid = true
}
