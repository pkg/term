package term

import (
	"syscall"
	"unsafe"
)

type attr syscall.Termios

func (t *Term) tcgetattr() (attr, error) {
	var a attr
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TIOCGETA, uintptr(unsafe.Pointer(&a)), 0, 0, 0); e != 0 {
		return a, e
	}
	return a, nil
}

func (t *Term) tcsetattr(a attr) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TIOCSETA, uintptr(unsafe.Pointer(&a)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

func (a *attr) setSpeed(baud int) {
	var rates = map[int]uint64{
		50:     syscall.B50,
		75:     syscall.B75,
		110:    syscall.B110,
		134:    syscall.B134,
		150:    syscall.B150,
		200:    syscall.B200,
		300:    syscall.B300,
		600:    syscall.B600,
		1200:   syscall.B1200,
		1800:   syscall.B1800,
		2400:   syscall.B2400,
		4800:   syscall.B4800,
		9600:   syscall.B9600,
		19200:  syscall.B19200,
		38400:  syscall.B38400,
		57600:  syscall.B57600,
		115200: syscall.B115200,
		230400: syscall.B230400,
	}

	rate := rates[baud]
	if rate == 0 {
		return
	}
	(*syscall.Termios)(a).Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	(*syscall.Termios)(a).Ispeed = rate
	(*syscall.Termios)(a).Ospeed = rate
}
