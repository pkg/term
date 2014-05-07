package termios

import (
	"os"
	"syscall"
	"testing"
)

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
	if err := Tcsetattr(f.Fd(), TCSANOW, &termios); err != nil {
		t.Fatal(err)
	}
}

func TestTcflush(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := Tcflush(f.Fd(), TCIOFLUSH); err != nil {
		t.Fatal(err)
	}
}
