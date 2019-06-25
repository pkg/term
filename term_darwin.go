package term

import "syscall"

type attr syscall.Termios

func (a *attr) getSpeed() (int, error) {
	// We generally only care about ospeed, since that's what would
	// be used for padding characters, for example.

	switch a.Ospeed {
	case syscall.B50:
		return 50, nil
	case syscall.B75:
		return 75, nil
	case syscall.B110:
		return 110, nil
	case syscall.B134:
		return 134, nil
	case syscall.B150:
		return 150, nil
	case syscall.B200:
		return 200, nil
	case syscall.B300:
		return 300, nil
	case syscall.B600:
		return 600, nil
	case syscall.B1200:
		return 1200, nil
	case syscall.B1800:
		return 1800, nil
	case syscall.B2400:
		return 2400, nil
	case syscall.B4800:
		return 4800, nil
	case syscall.B9600:
		return 9600, nil
	case syscall.B19200:
		return 19200, nil
	case syscall.B38400:
		return 38400, nil
	case syscall.B57600:
		return 57600, nil
	case syscall.B115200:
		return 115200, nil
	case syscall.B230400:
		return 230400, nil
	default:
		return 0, syscall.EINVAL
	}
}

func (a *attr) setSpeed(baud int) error {
	var rate uint64
	switch baud {
	case 50:
		rate = syscall.B50
	case 75:
		rate = syscall.B75
	case 110:
		rate = syscall.B110
	case 134:
		rate = syscall.B134
	case 150:
		rate = syscall.B150
	case 200:
		rate = syscall.B200
	case 300:
		rate = syscall.B300
	case 600:
		rate = syscall.B600
	case 1200:
		rate = syscall.B1200
	case 1800:
		rate = syscall.B1800
	case 2400:
		rate = syscall.B2400
	case 4800:
		rate = syscall.B4800
	case 9600:
		rate = syscall.B9600
	case 19200:
		rate = syscall.B19200
	case 38400:
		rate = syscall.B38400
	case 57600:
		rate = syscall.B57600
	case 115200:
		rate = syscall.B115200
	case 230400:
		rate = syscall.B230400
	default:
		return syscall.EINVAL
	}
	(*syscall.Termios)(a).Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	(*syscall.Termios)(a).Ispeed = rate
	(*syscall.Termios)(a).Ospeed = rate
	return nil
}
