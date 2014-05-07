package termios

import (
	"syscall"
	"unsafe"
)

// Tcgetattr gets the current serial port settings.
func Tcgetattr(fd uintptr, argp *syscall.Termios) error {
	return ioctl(fd, syscall.TCGETS, uintptr(unsafe.Pointer(argp)))
}

// Tcsetattr sets the current serial port settings.
func Tcsetattr(fd, action uintptr, argp *syscall.Termios) error {
	var request uintptr
	switch action {
	case TCSANOW:
		request = TCSETS
	case TCSADRAIN:
		request = TCSETSW
	case TCSAFLUSH:
		request = TCSETSF
	default:
		return syscall.EINVAL
	}
	return ioctl(fd, request, uintptr(unsafe.Pointer(argp)))
}

// Tcsendbreak transmits a continuous stream of zero-valued bits for a specific
// duration, if the terminal is using asynchronous serial data transmission. If
// duration is zero, it transmits zero-valued bits for at least 0.25 seconds, and not more that 0.5 seconds.
// If duration is not zero, it sends zero-valued bits for some
// implementation-defined length of time.
func Tcsendbreak(fd, duration uintptr) error {
	return ioctl(fd, TCSBRKP, duration)
}

// Tcflush discards data written to the object referred to by fd but not transmitted, or data received but not read, depending on the value of selector.
func Tcflush(fd, selector uintptr) error {
	return ioctl(fd, TCFLSH, selector)
}
