// Package term manages POSIX terminals. As POSIX terminals are connected to,
// or emulate, a UART, this package also provides control over the various
// UART and serial line parameters.
package term

import (
	"io"
	"os"
	"syscall"

	"github.com/pkg/term/termios"
)

// Term represents an asynchronous communications port.
type Term struct {
	name string
	fd   int
}

// Open opens an asynchronous communications port.
func Open(name string) (*Term, error) {
	fd, e := syscall.Open(name, syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if e != nil {
		return nil, &os.PathError{"open", name, e}
	}
	return &Term{name: name, fd: fd}, nil
}

// Read reads up to len(b) bytes from the terminal. It returns the number of
// bytes read and an error, if any. EOF is signaled by a zero count with
// err set to io.EOF.
func (t *Term) Read(b []byte) (int, error) {
	n, e := syscall.Read(t.fd, b)
	if n < 0 {
		n = 0
	}
	if n == 0 && len(b) > 0 && e == nil {
		return 0, io.EOF
	}
	if e != nil {
		return n, &os.PathError{"read", t.name, e}
	}
	return n, nil
}

// Write writes len(b) bytes to the terminal. It returns the number of bytes
// written and an error, if any. Write returns a non-nil error when n !=
// len(b).
func (t *Term) Write(b []byte) (int, error) {
	n, e := syscall.Write(t.fd, b)
	if n < 0 {
		n = 0
	}
	if n != len(b) {
		return n, io.ErrShortWrite
	}
	if e != nil {
		return n, &os.PathError{"write", t.name, e}
	}
	return n, nil
}

// Close closes the device and releases any associated resources.
func (t *Term) Close() error {
	err := syscall.Close(t.fd)
	t.fd = -1
	return err
}

// SetSpeed sets the receive and transmit baud rates.
func (t *Term) SetSpeed(baud int) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	a.setSpeed(baud)
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

// Flush flushes both data received but not read, and data written but not transmitted.
func (t *Term) Flush() error {
	return termios.Tcflush(uintptr(t.fd), termios.TCIOFLUSH)
}

// SendBreak sends a break signal.
func (t *Term) SendBreak() error {
	return termios.Tcsendbreak(uintptr(t.fd), 0)
}

// SetDTR sets the DTR (data terminal ready) signal.
func (t *Term) SetDTR(v bool) error {
	bits := syscall.TIOCM_DTR
	if v {
		return termios.Tiocmbis(uintptr(t.fd), &bits)
	} else {
		return termios.Tiocmbic(uintptr(t.fd), &bits)
	}
}

// DTR returns the state of the DTR (data terminal ready) signal.
func (t *Term) DTR() (bool, error) {
	var status int
	err := termios.Tiocmget(uintptr(t.fd), &status)
	return status&syscall.TIOCM_DTR == syscall.TIOCM_DTR, err
}
