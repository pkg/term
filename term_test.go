package term

import "testing"
import "flag"

var tty = flag.String("tty", "/dev/ttyUSB0", "the terminal device to use for testing")

func TestTermSpeed(t *testing.T) {
	tt, err := Open(*tty)
	if err != nil {
		t.Fatal(err)
	}
	defer tt.Close()
	t.Log(tt.Speed())
}

func TestTermSetSpeed(t *testing.T) {
	tt, err := Open(*tty)
	if err != nil {
		t.Fatal(err)
	}
	defer tt.Close()
	if err := tt.SetSpeed(57600); err != nil {
		t.Fatal(err)
	}
	t.Log(tt.Speed())
}
