// +build darwin freebsd openbsd netbsd

package termios

import (
	"os"
	"syscall"
	"testing"
)

func TestTcflush(t *testing.T) {
	f, err := os.OpenFile("/dev/tty", syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := Tcflush(f.Fd(), syscall.TCIOFLUSH); err != nil {
		t.Fatal(err)
	}
}
