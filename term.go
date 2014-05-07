// Package term manages POSIX terminals. As POSIX terminals are connected to,
// or emulate, a UART, this package also provides control over the various
// UART and serial line parameters.
package term

import (
	"io"
	"os"
	"syscall"
	"unsafe"
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
	attr, err := t.tcgetattr()
	if err != nil {
		return err
	}
	attr.setSpeed(baud)
	return t.tcsetattr(attr)
}

// Flush flushes both data received but not read, and data written but not transmitted.
func (t *Term) Flush() error {
	const TCFLSH = 0x540B
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), TCFLSH, syscall.TCIOFLUSH, 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

// SendBreak sends a break signal.
func (t *Term) SendBreak() error {
	const TCSBRK = 0x5409 // not POSIX TCSBRKP
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), TCSBRK, 0, 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

// Status represents the current "MODEM" status bits, which consist of all of the RS-232 signal lines except RXD and TXD.
type Status int

// SetDTR sets the DTR (data terminal ready) RS-232 control signal.
func (s *Status) SetDTR(v bool) {
	if v {
		(*s) |= syscall.TIOCM_DTR
	} else {
		(*s) &= ^syscall.TIOCM_DTR
	}
}

// Status returns the state of the "MODEM" bits.
func (t *Term) Status() (Status, error) {
	var status int
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TIOCMGET, uintptr(unsafe.Pointer(&status)), 0, 0, 0); e != 0 {
		return 0, e
	}
	return Status(status), nil
}

// SetStatus sets the state of the "MODEM" bits.
func (t *Term) SetStatus(status Status) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TIOCMSET, uintptr(unsafe.Pointer(&status)), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}
