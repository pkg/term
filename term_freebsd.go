package term

import "syscall"

type attr syscall.Termios

func (a *attr) setSpeed(baud int) {
	var rates = map[int]uint32{
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
		460800: syscall.B460800,
		921600: syscall.B921600,
	}

	rate := rates[baud]
	if rate == 0 {
		return
	}
	(*syscall.Termios)(a).Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	(*syscall.Termios)(a).Ispeed = rate
	(*syscall.Termios)(a).Ospeed = rate
}
