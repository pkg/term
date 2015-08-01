package termios

import (
	"fmt"
	"syscall"
	"unsafe"
)

func ptsname(fd uintptr) (string, error) {
	var n uintptr
	err := ioctl(fd, syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/pts/%d", n), nil
}

func grantpt(fd uintptr) error {
	var n uintptr
	return ioctl(fd, syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
}

func unlockpt(fd uintptr) error {
	var n uintptr
	return ioctl(fd, syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&n)))
}
