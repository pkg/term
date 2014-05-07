// Package termios implements the low level termios(3) terminal line discipline facilities.
//
// For a higher level interface please use the github.com/pkg/term package.
package termios

import (
	"syscall"
	"unsafe"
)

func ioctl(fd, request, argp uintptr) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, request, argp, 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

// Tiocmget returns the state of the MODEM bits.
func Tiocmget(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMGET, uintptr(unsafe.Pointer(status)))
}

// Tiocmset sets the state of the MODEM bits.
func Tiocmset(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMSET, uintptr(unsafe.Pointer(status)))
}

// Tiocmbis sets the indicated modem bits.
func Tiocmbis(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMBIS, uintptr(unsafe.Pointer(status)))
}

// Tiocmbic clears the indicated modem bits.
func Tiocmbic(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMBIC, uintptr(unsafe.Pointer(status)))
}
