package px

import (
	"syscall"
)

func setNewProcessGroupAttr(attr *syscall.SysProcAttr) {
	attr.CreationFlags = syscall.CREATE_NEW_PROCESS_GROUP
}
