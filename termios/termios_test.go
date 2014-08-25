package termios

import (
	"flag"
	"os"
	"syscall"
	"testing"
)

var dev = flag.String("device", "/dev/tty", "device to use")

func TestTcgetattr(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
}

func TestTcsetattr(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
	for _, opt := range []uintptr{TCSANOW, TCSADRAIN, TCSAFLUSH} {
		if err := Tcsetattr(f.Fd(), opt, &termios); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTcsendbreak(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	if err := Tcsendbreak(f.Fd(), 0); err != nil {
		t.Fatal(err)
	}
}

func TestTcdrain(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	if err := Tcdrain(f.Fd()); err != nil {
		t.Fatal(err)
	}
}

func TestTiocmget(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var status int
	if err := Tiocmget(f.Fd(), &status); err != nil {
		t.Fatal(err)
	}
}

func TestTiocmset(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var status int
	if err := Tiocmget(f.Fd(), &status); err != nil {
		t.Fatal(err)
	}
	if err := Tiocmset(f.Fd(), &status); err != nil {
		t.Fatal(err)
	}
}

func TestTiocmbis(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	status := 0
	if err := Tiocmbis(f.Fd(), &status); err != nil {
		t.Fatal(err)
	}
}

func TestTiocmbic(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	status := 0
	if err := Tiocmbic(f.Fd(), &status); err != nil {
		t.Fatal(err)
	}
}

func TestTiocinq(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var inq int
	if err := Tiocinq(f.Fd(), &inq); err != nil {
		t.Fatal(err)
	}
	if inq != 0 {
		t.Fatal("Expected 0 bytes, got %v", inq)
	}
}

func TestTiocoutq(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var inq int
	if err := Tiocoutq(f.Fd(), &inq); err != nil {
		t.Fatal(err)
	}
	if inq != 0 {
		t.Fatal("Expected 0 bytes, got %v", inq)
	}
}

func TestCfgetispeed(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
	if baud := Cfgetispeed(&termios); baud == 0 {
		t.Fatalf("Cfgetispeed: expected > 0, got %v", baud)
	}
}

func TestCfgetospeed(t *testing.T) {
	f := opendev(t)
	defer f.Close()

	var termios syscall.Termios
	if err := Tcgetattr(f.Fd(), &termios); err != nil {
		t.Fatal(err)
	}
	if baud := Cfgetospeed(&termios); baud == 0 {
		t.Fatalf("Cfgetospeed: expected > 0, got %v", baud)
	}
}

func opendev(t *testing.T) *os.File {
	f, err := os.OpenFile(*dev, syscall.O_NOCTTY|syscall.O_CLOEXEC|syscall.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	return f
}
