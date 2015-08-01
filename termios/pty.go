package termios

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// Pty returns a UNIX 98 pseudoterminal device.
// Pty returns a pair of fds representing the master and slave pair.
func Pty() (*os.File, *os.File, error) {
	open := func(path string) (uintptr, error) {
		fd, err := syscall.Open(path, syscall.O_NOCTTY|syscall.O_RDWR|syscall.O_CLOEXEC, 0666)
		if err != nil {
			return 0, fmt.Errorf("unable to open %q: %v", path, err)
		}
		return uintptr(fd), nil
	}
	ptm, err := open("/dev/ptmx")
	if err != nil {
		return nil, nil, err
	}
	unlock := uintptr(0)
	if err := ioctl(uintptr(ptm), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&unlock))); err != nil {
		return nil, nil, fmt.Errorf("TIOCSPLCK: %v", err)
	}
	var pty_nam uintptr
	if err := ioctl(uintptr(ptm), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&pty_nam))); err != nil {
		return nil, nil, fmt.Errorf("TIOCGPTN: %v", err)
	}
	pts, err := open(fmt.Sprintf("/dev/pts/%d", pty_nam))
	if err != nil {
		return nil, nil, err
	}
	return os.NewFile(uintptr(ptm), "ptm"), os.NewFile(uintptr(pts), "pts"), nil
}
