package term

import (
	"syscall"
	"unsafe"
)

func (t *Term) tcgetattr() (*syscall.Termios, error) {
	var attr syscall.Termios
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCGETS, uintptr(unsafe.Pointer(&attr)), 0, 0, 0); e != 0 {
		return nil, e
	}
	return &attr, nil
}

func (t *Term) tcsetattr(attr *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCSETS, uintptr(unsafe.Pointer(attr)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

func cfsetspeed(attr *syscall.Termios, baud int) {
	var rates = map[int]uint32{
		50:      syscall.B50,
		75:      syscall.B75,
		110:     syscall.B110,
		134:     syscall.B134,
		150:     syscall.B150,
		200:     syscall.B200,
		300:     syscall.B300,
		600:     syscall.B600,
		1200:    syscall.B1200,
		1800:    syscall.B1800,
		2400:    syscall.B2400,
		4800:    syscall.B4800,
		9600:    syscall.B9600,
		19200:   syscall.B19200,
		38400:   syscall.B38400,
		57600:   syscall.B57600,
		115200:  syscall.B115200,
		230400:  syscall.B230400,
		460800:  syscall.B460800,
		500000:  syscall.B500000,
		576000:  syscall.B576000,
		921600:  syscall.B921600,
		1000000: syscall.B1000000,
		1152000: syscall.B1152000,
		1500000: syscall.B1500000,
		2000000: syscall.B2000000,
		2500000: syscall.B2500000,
		3000000: syscall.B3000000,
		3500000: syscall.B3500000,
		4000000: syscall.B4000000,
	}

	rate := rates[baud]
	if rate == 0 {
		return
	}
	attr.Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	attr.Ispeed = rate
	attr.Ospeed = rate
}
