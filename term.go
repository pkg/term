// Package term manages POSIX terminals. As POSIX terminals are connected to, or emulate,
// a UART, this package provides control over the various UART and serial line parameters.
package term

import (
	"os"
	"syscall"
	"unsafe"
)

// Term represents an asynchronous communications port.
type Term struct {
	fd int
}

// Open opens an asynchronous communications port.
func Open(name string) (*Term, error) {
	fd, e := syscall.Open(name, syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if e != nil {
		return nil, &os.PathError{"open", name, e}
	}
	return &Term{fd: fd}, nil
}

// Read reads up to len(b) bytes from the terminal. It returns the number of
// bytes read and an error, if any. EOF is signaled by a zero count with
// err set to io.EOF.
func (t *Term) Read(b []byte) (int, error) {
	return syscall.Read(t.fd, b)
}

// Write writes len(b) bytes to the terminal. It returns the number of bytes
// written and an error, if any. Write returns a non-nil error when n !=
// len(b).
func (t *Term) Write(b []byte) (int, error) {
	return syscall.Write(t.fd, b)
}

// Close releases any associated resources.
func (t *Term) Close() error {
	err := syscall.Close(t.fd)
	t.fd = -1
	return err
}

// Speed returns the current input and output baud rates for device.
func (t *Term) Speed() (int, int, error) {
	termios, err := t.tcgetattr()
	if err != nil {
		return -1, -1, err
	}
	return int(termios.Ispeed), int(termios.Ospeed), nil
}

// SetSpeed sets the receive and transmit baud rates.
func (t *Term) SetSpeed(baud int) error {
	attr, err := t.tcgetattr()
	if err != nil {
		return err
	}
	cfsetspeed(attr, baud)
	return t.tcsetattr(attr)
}

func (t *Term) tcgetattr() (*syscall.Termios, error) {
	var termios syscall.Termios
	if err := t.ioctl(syscall.TCGETS, &termios); err != nil {
		return nil, err
	}
	return &termios, nil
}

func (t *Term) tcsetattr(attr *syscall.Termios) error {
	return t.ioctl(syscall.TCSETS, attr)
}

func (t *Term) ioctl(op uintptr, p *syscall.Termios) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), op, uintptr(unsafe.Pointer(p)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

func cfsetspeed(attr *syscall.Termios, baud int) {
	var bauds = map[int]uint32{
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

	rate := bauds[baud]
	if rate == 0 {
		return
	}
	attr.Cflag = syscall.CS8 | syscall.CREAD | syscall.CLOCAL | rate
	attr.Ispeed = rate
	attr.Ospeed = rate
}
