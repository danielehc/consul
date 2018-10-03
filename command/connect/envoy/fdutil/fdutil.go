// package fdutil provides some utilities to manipulate file descriptors
package fdutil

import (
	"golang.org/x/sys/unix"
)

// IsCloseOnExec checks the provided file descriptor to see if the CLOEXEC flag
// is set.
func IsCloseOnExec(fd uintptr) (bool, error) {
	flags, err := getFdFlags(fd)
	if err != nil {
		return false, err
	}
	return flags&unix.FD_CLOEXEC != 0, nil
}

// SetCloseOnExec sets or unsets the CLOEXEC flag on the provided file descriptor
// depending upon the value of the enabled arg.
func SetCloseOnExec(fd uintptr, enabled bool) error {
	flags, err := getFdFlags(fd)
	if err != nil {
		return err
	}

	newFlags := flags
	if enabled {
		newFlags |= unix.FD_CLOEXEC
	} else {
		newFlags &= ^unix.FD_CLOEXEC
	}

	if newFlags == flags {
		return nil // noop
	}

	_, err = unix.FcntlInt(fd, unix.F_SETFD, newFlags)
	return err
}

func getFdFlags(fd uintptr) (int, error) {
	return unix.FcntlInt(fd, unix.F_GETFD, 0)
}
