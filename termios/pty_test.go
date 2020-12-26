package termios

import (
	"testing"
)

func TestPtsname(t *testing.T) {
	fd, err := open_pty_master()
	if err != nil {
		t.Fatal(err)
	}

	name, err := Ptsname(fd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}
