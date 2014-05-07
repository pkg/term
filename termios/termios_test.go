package termios

import (
	"os"
	"syscall"
	"testing"
)

func TestTcgetattr(t *testing.T) {
	f, err := os.OpenFile("/dev/pty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
}
