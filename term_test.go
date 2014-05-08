package term

import "testing"
import "flag"

var dev = flag.String("device", "/dev/tty", "device to use")

func TestTermSetSpeed(t *testing.T) {
	tt := opendev(t)
	defer tt.Close()
	if err := tt.SetSpeed(57600); err != nil {
		t.Fatal(err)
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
	tt, err := Open(*dev)
	if err != nil {
		t.Fatal(err)
	}
	return tt
}
