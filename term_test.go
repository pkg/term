// +build !windows

package term

import (
	"testing"

	"time"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

// assert that Term implements the same method set across
// all supported platforms
var _ interface {
	Available() (int, error)
	Buffered() (int, error)
	CTS() (bool, error)
	Close() error
	DCD() (bool, error)
	DSR() (bool, error)
	DTR() (bool, error)
	Flush() error
	RI() (bool, error)
	RTS() (bool, error)
	Read(b []byte) (int, error)
	Restore() error
	SendBreak() error
	SetCbreak() error
	SetDTR(v bool) error
	SetOption(options ...func(*Term) error) error
	SetRTS(v bool) error
	SetRaw() error
	SetSpeed(baud int) error
	GetSpeed() (int, error)
	Write(b []byte) (int, error)
} = new(Term)

func TestTermSetCbreak(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetCbreak(); err != nil {
		t.Fatal(err)
	}
}

func TestTermSetRaw(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetRaw(); err != nil {
		t.Fatal(err)
	}
}

func TestTermSetSpeed(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetSpeed(57600); err != nil {
		t.Fatal(err)
	}

	if spd, err := tt.GetSpeed(); err != nil {
		t.Fatal(err)
	} else if spd != 57600 {
		t.Errorf("speed mismatch %d != 57600", spd)
	}
}

func TestTermSetInvalidSpeed(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetSpeed(12345); err != unix.EINVAL {
		t.Fatal(err)
	}
}

func TestTermSetReadTimeout(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetReadTimeout(1 * time.Second); err != nil {
		t.Fatal(err)
	}
}

func TestTermSetFlowControl(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()

	kinds := []int{XONXOFF, NONE, HARDWARE, NONE, XONXOFF, HARDWARE, NONE}

	for _, kind := range kinds {
		if err := tt.SetFlowControl(kind); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTermRestore(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.Restore(); err != nil {
		t.Fatal(err)
	}
}

func opendev(t *testing.T) *Term {
	_, pts, err := termios.Pty()
	if err != nil {
		t.Fatal(err)
	}
	term, err := Open(pts.Name())
	if err != nil {
		t.Fatal(err)
	}
	pts.Close()
	return term
}
