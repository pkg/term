// +build darwin freebsd openbsd netbsd

package termios

import (
	"syscall"
	"unsafe"
)

// Tcgetattr gets the current serial port settings.
func Tcgetattr(fd uintptr, argp *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(argp)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

// Tcsetattr sets the current serial port settings.
func Tcsetattr(fd uintptr, action int, argp *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCSETA, uintptr(unsafe.Pointer(argp)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}
