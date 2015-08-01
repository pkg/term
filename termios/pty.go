package termios

import (
	"fmt"
	"os"
	"syscall"
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

	sname, err := ptsname(ptm)
	if err != nil {
		return nil, nil, err
	}

	err = grantpt(ptm)
	if err != nil {
		return nil, nil, err
	}

	err = unlockpt(ptm)
	if err != nil {
		return nil, nil, err
	}

	pts, err := open(sname)
	if err != nil {
		return nil, nil, err
	}
	return os.NewFile(uintptr(ptm), "ptm"), os.NewFile(uintptr(pts), sname), nil
}
