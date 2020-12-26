// +build !windows

package termios

import (
	"os"
	"runtime"
	"testing"

	"golang.org/x/sys/unix"
)

func TestTcgetattr(t *testing.T) {
	f := opendev(t)

	if _, err := Tcgetattr(f.Fd()); err != nil {
		t.Fatal(err)
	}
}

func TestTcsetattr(t *testing.T) {
	f := opendev(t)

	termios, err := Tcgetattr(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	for _, opt := range []uintptr{TCSANOW, TCSADRAIN, TCSAFLUSH} {
		if err := Tcsetattr(f.Fd(), opt, termios); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTcsendbreak(t *testing.T) {
	f := opendev(t)

	if err := Tcsendbreak(f.Fd(), 0); err != nil {
		t.Fatal(err)
	}
}

func TestTcdrain(t *testing.T) {
	f := opendev(t)

	if err := Tcdrain(f.Fd()); err != nil {
		t.Fatal(err)
	}
}

func TestTiocmget(t *testing.T) {
	f := opendev(t)

	if _, err := Tiocmget(f.Fd()); err != nil {
		checktty(t, err)
		t.Fatal(err)
	}
}

func TestTiocmset(t *testing.T) {
	f := opendev(t)

	status, err := Tiocmget(f.Fd())
	if err != nil {
		checktty(t, err)
		t.Fatal(err)
	}
	if err := Tiocmset(f.Fd(), status); err != nil {
		checktty(t, err)
		t.Fatal(err)
	}
}

func TestTiocmbis(t *testing.T) {
	f := opendev(t)

	if err := Tiocmbis(f.Fd(), 0); err != nil {
		checktty(t, err)
		t.Fatal(err)
	}
}

func TestTiocmbic(t *testing.T) {
	f := opendev(t)

	if err := Tiocmbic(f.Fd(), 0); err != nil {
		checktty(t, err)
		t.Fatal(err)
	}
}

func TestTiocinq(t *testing.T) {
	f := opendev(t)

	inq, err := Tiocinq(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	if inq != 0 {
		t.Fatalf("Expected 0 bytes, got %v", inq)
	}
}

func TestTiocoutq(t *testing.T) {
	f := opendev(t)

	inq, err := Tiocoutq(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	if inq != 0 {
		t.Fatalf("Expected 0 bytes, got %v", inq)
	}
}

func TestCfgetispeed(t *testing.T) {
	f := opendev(t)

	termios, err := Tcgetattr(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	if baud := Cfgetispeed(termios); baud == 0 && runtime.GOOS != "linux" {
		t.Fatalf("Cfgetispeed: expected > 0, got %v", baud)
	}
}

func TestCfgetospeed(t *testing.T) {
	f := opendev(t)

	termios, err := Tcgetattr(f.Fd())
	if err != nil {
		t.Fatal(err)
	}
	if baud := Cfgetospeed(termios); baud == 0 && runtime.GOOS != "linux" {
		t.Fatalf("Cfgetospeed: expected > 0, got %v", baud)
	}
}

func opendev(t *testing.T) *os.File {
	_, pts, err := Pty()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pts.Close()
	})
	return pts
}

func checktty(t *testing.T, err error) {
	t.Helper()
	// some ioctls fail against char devices if they do not
	// support a particular feature
	if (runtime.GOOS == "darwin" && err == unix.ENOTTY) || (runtime.GOOS == "linux" && err == unix.EINVAL) {
		t.Skip(err)
	}
}
