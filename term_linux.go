package term

import "syscall"

type attr syscall.Termios

const (
	// CBaudMask is the logical of CBAUD and CBAUDEX, except
	// that those values were not exposed via the syscall
	// package.  Many of these values will be redundant, but
	// this long definition ensures we are portable if some
	// architecture defines different values for them (unlikely).
	CBaudMask = syscall.B50 |
		syscall.B75 |
		syscall.B110 |
		syscall.B134 |
		syscall.B150 |
		syscall.B200 |
		syscall.B300 |
		syscall.B600 |
		syscall.B1200 |
		syscall.B1800 |
		syscall.B2400 |
		syscall.B4800 |
		syscall.B9600 |
		syscall.B19200 |
		syscall.B38400 |
		syscall.B57600 |
		syscall.B115200 |
		syscall.B230400 |
		syscall.B460800 |
		syscall.B500000 |
		syscall.B576000 |
		syscall.B921600 |
		syscall.B1000000 |
		syscall.B1152000 |
		syscall.B1500000 |
		syscall.B2000000 |
		syscall.B2500000 |
		syscall.B3000000 |
		syscall.B3500000 |
		syscall.B4000000
)

func (a *attr) getSpeed() (int, error) {
	switch a.Cflag & CBaudMask {
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
	case syscall.B460800:
		return 460800, nil
	case syscall.B500000:
		return 500000, nil
	case syscall.B576000:
		return 576000, nil
	case syscall.B921600:
		return 921600, nil
	case syscall.B1000000:
		return 1000000, nil
	case syscall.B1152000:
		return 1152000, nil
	case syscall.B1500000:
		return 1500000, nil
	case syscall.B2000000:
		return 2000000, nil
	case syscall.B2500000:
		return 2500000, nil
	case syscall.B3000000:
		return 3000000, nil
	case syscall.B3500000:
		return 3500000, nil
	case syscall.B4000000:
		return 4000000, nil
	default:
		return 0, syscall.EINVAL
	}
}

func (a *attr) setSpeed(baud int) error {
	var rate uint32
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
	case 460800:
		rate = syscall.B460800
	case 500000:
		rate = syscall.B500000
	case 576000:
		rate = syscall.B576000
	case 921600:
		rate = syscall.B921600
	case 1000000:
		rate = syscall.B1000000
	case 1152000:
		rate = syscall.B1152000
	case 1500000:
		rate = syscall.B1500000
	case 2000000:
		rate = syscall.B2000000
	case 2500000:
		rate = syscall.B2500000
	case 3000000:
		rate = syscall.B3000000
	case 3500000:
		rate = syscall.B3500000
	case 4000000:
		rate = syscall.B4000000
	default:
		return syscall.EINVAL
	}
	(*syscall.Termios)(a).Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	(*syscall.Termios)(a).Ispeed = rate
	(*syscall.Termios)(a).Ospeed = rate
	return nil
}
