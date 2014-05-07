// Package termios implements the low level termios(3) terminal line discipline facilities.
//
// For a higher level interface please use the github.com/pkg/term package.
package termios

import (
	"syscall"
)

func ioctl(fd, request, argp uintptr) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, request, argp, 0, 0, 0); e != 0 {
		return e
	}
	return nil
}
