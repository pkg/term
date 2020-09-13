package term

import "golang.org/x/sys/unix"

type attr unix.Termios

func (a *attr) setSpeed(baud int) error {
	var rate uint32
	switch baud {
	case 50:
		rate = unix.B50
	case 75:
		rate = unix.B75
	case 110:
		rate = unix.B110
	case 134:
		rate = unix.B134
	case 150:
		rate = unix.B150
	case 200:
		rate = unix.B200
	case 300:
		rate = unix.B300
	case 600:
		rate = unix.B600
	case 1200:
		rate = unix.B1200
	case 1800:
		rate = unix.B1800
	case 2400:
		rate = unix.B2400
	case 4800:
		rate = unix.B4800
	case 9600:
		rate = unix.B9600
	case 19200:
		rate = unix.B19200
	case 38400:
		rate = unix.B38400
	case 57600:
		rate = unix.B57600
	case 115200:
		rate = unix.B115200
	case 230400:
		rate = unix.B230400
	case 460800:
		rate = unix.B460800
	case 500000:
		rate = unix.B500000
	case 576000:
		rate = unix.B576000
	case 921600:
		rate = unix.B921600
	case 1000000:
		rate = unix.B1000000
	case 1152000:
		rate = unix.B1152000
	case 1500000:
		rate = unix.B1500000
	case 2000000:
		rate = unix.B2000000
	case 2500000:
		rate = unix.B2500000
	case 3000000:
		rate = unix.B3000000
	case 3500000:
		rate = unix.B3500000
	case 4000000:
		rate = unix.B4000000
	default:
		return unix.EINVAL
	}
	(*unix.Termios)(a).Cflag = unix.CS8 | unix.CREAD | unix.CLOCAL | rate
	(*unix.Termios)(a).Ispeed = rate
	(*unix.Termios)(a).Ospeed = rate
	return nil
}
