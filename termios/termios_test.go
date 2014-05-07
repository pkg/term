package termios

import (
	"os"
	"syscall"
	"testing"
)

func TestTcgetattr(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
}

func TestTcsetattr(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
	for _, opt := range []int{TCSANOW, TCSADRAIN, TCSAFLUSH} {
		if err := Tcsetattr(f.Fd(), opt, &termios); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTcsendbreak(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := Tcsendbreak(f.Fd(), 0); err != nil {
		t.Fatal(err)
	}
}

func TestTcdrain(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := Tcdrain(f.Fd()); err != nil {
		t.Fatal(err)
	}
}
