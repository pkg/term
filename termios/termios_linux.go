// +build darwin freebsd openbsd netbsd

package termios

import (
	"syscall"
	"unsafe"
)

// Tcgetattr gets the current serial port settings.
func Tcgetattr(fd uintptr, argp *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TCGETS, uintptr(unsafe.Pointer(argp)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

// Tcsetattr sets the current serial port settings.
func Tcsetattr(fd, action uintptr, argp *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, action, uintptr(unsafe.Pointer(argp)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}
