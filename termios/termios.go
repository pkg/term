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

// Cfgetispeed returns the input baud rate stored in the termios structure.
func Cfgetispeed(attr *syscall.Termios) uint32 { return attr.Ispeed }

// Cfgetospeed returns the output baud rate stored in the termios structure.
func Cfgetospeed(attr *syscall.Termios) uint32 { return attr.Ospeed }
