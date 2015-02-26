// +build !windows

package term

import (
	"os"
	"syscall"

	"github.com/pkg/term/termios"
)

// Term represents an asynchronous communications port.
type Term struct {
	name string
	fd   int
	orig syscall.Termios // original state of the terminal, see Open and Restore
}

// Open opens an asynchronous communications port.
func Open(name string, options ...func(*Term) error) (*Term, error) {
	fd, e := syscall.Open(name, syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if e != nil {
		return nil, &os.PathError{"open", name, e}
	}
	t := Term{name: name, fd: fd}
	if err := termios.Tcgetattr(uintptr(t.fd), &t.orig); err != nil {
		return nil, err
	}
	return &t, t.SetOption(options...)
}

// SetCbreak sets cbreak mode.
func (t *Term) SetCbreak() error {
	return t.SetOption(CBreakMode)
}

// CBreakMode places the terminal into cbreak mode.
func CBreakMode(t *Term) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	termios.Cfmakecbreak((*syscall.Termios)(&a))
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

// SetRaw sets raw mode.
func (t *Term) SetRaw() error {
	return t.SetOption(RawMode)
}

// RawMode places the terminal into raw mode.
func RawMode(t *Term) error {
	var a attr
	if err := termios.Tcgetattr(uintptr(t.fd), (*syscall.Termios)(&a)); err != nil {
		return err
	}
	termios.Cfmakeraw((*syscall.Termios)(&a))
	return termios.Tcsetattr(uintptr(t.fd), termios.TCSANOW, (*syscall.Termios)(&a))
}

// Speed sets the baud rate option for the terminal.
func Speed(baud int) func(*Term) error {
	return func(t *Term) error {
		return t.setSpeed(baud)
	}
}

// SetSpeed sets the receive and transmit baud rates.
func (t *Term) SetSpeed(baud int) error {
	return t.SetOption(Speed(baud))
}

func (t *Term) setSpeed(baud int) error {
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

// SetRTS sets the RTS (data terminal ready) signal.
func (t *Term) SetRTS(v bool) error {
	bits := syscall.TIOCM_RTS
	if v {
		return termios.Tiocmbis(uintptr(t.fd), &bits)
	} else {
		return termios.Tiocmbic(uintptr(t.fd), &bits)
	}
}

// RTS returns the state of the RTS (data terminal ready) signal.
func (t *Term) RTS() (bool, error) {
	var status int
	err := termios.Tiocmget(uintptr(t.fd), &status)
	return status&syscall.TIOCM_RTS == syscall.TIOCM_RTS, err
}

// Restore restores the state of the terminal captured at the point that
// the terminal was originally opened.
func (t *Term) Restore() error {
	return termios.Tcsetattr(uintptr(t.fd), termios.TCIOFLUSH, &t.orig)
}

// Close closes the device and releases any associated resources.
func (t *Term) Close() error {
	err := syscall.Close(t.fd)
	t.fd = -1
	return err
}
