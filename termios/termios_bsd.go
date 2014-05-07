// +build darwin freebsd openbsd netbsd

package termios

import (
	"syscall"
	"time"
	"unsafe"
)

const (
	FREAD     = 0x0001
	FWRITE    = 0x0002
	TCSANOW   = 0 /* make change immediate */
	TCSADRAIN = 1 /* drain output, then change */
	TCSAFLUSH = 2
)

// Tcgetattr gets the current serial port settings.
func Tcgetattr(fd uintptr, argp *syscall.Termios) error {
	return ioctl(fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(argp)))
}

// Tcsetattr sets the current serial port settings.
func Tcsetattr(fd uintptr, action int, argp *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCSETA, uintptr(unsafe.Pointer(argp)), 0, 0, 0); e != 0 {
		return e

	}
	return nil
}

// Tcsendbreak function transmits a continuous stream of zero-valued bits for
// four-tenths of a second to the terminal referenced by fildes. The duration
// parameter is ignored in this implementation.
func Tcsendbreak(fd, duration uintptr) error {
	if err := ioctl(fd, syscall.TIOCSBRK, 0); err != nil {
		return err
	}
	time.Sleep(4 / 10 * time.Second)
	return ioctl(fd, syscall.TIOCCBRK, 0)
}

// Tcdrain waits until all output written to the terminal referenced by fd has been transmitted to the terminal.
func Tcdrain(fd uintptr) error {
	return ioctl(fd, syscall.TIOCDRAIN, 0)
}

// Tcflush discards data written to the object referred to by fd but not transmitted, or data received but not read, depending on the value of which.
func Tcflush(fd, which uintptr) error {
	var com int
	switch which {
	case syscall.TCIFLUSH:
		com = FREAD
	case syscall.TCOFLUSH:
		com = FWRITE
	case syscall.TCIOFLUSH:
		com = FREAD | FWRITE
	default:
		return syscall.EINVAL
	}
	return ioctl(fd, syscall.TCIOFLUSH, uintptr(unsafe.Pointer(&com)))
}
