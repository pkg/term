// +build darwin freebsd openbsd netbsd dragonfly

package termios

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestTcflush(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	if err := Tcflush(f.Fd(), unix.TCIOFLUSH); err != nil {
		t.Fatal(err)
	}
}
