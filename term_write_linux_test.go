package term

import (
	"errors"
	"io"
	"testing"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func TestTermWriteHungUp(t *testing.T) {
	master, slave, err := termios.Pty()
	if err != nil {
		t.Fatal(err)
	}
	defer master.Close()
	defer slave.Close()
	terminal, err := Open(slave.Name(), RawMode)
	if err != nil {
		t.Fatal(err)
	}
	defer terminal.Close()
	if err := master.Close(); err != nil {
		t.Fatal(err)
	}
	n, err := terminal.Write([]byte("disconnected terminal"))
	checkWriteError(t, terminal.name, n, err, unix.EIO)
}

func TestTermWriteNonblocking(t *testing.T) {
	master, slave, err := termios.Pty()
	if err != nil {
		t.Fatal(err)
	}
	defer master.Close()
	defer slave.Close()
	terminal, err := Open(slave.Name(), RawMode)
	if err != nil {
		t.Fatal(err)
	}
	defer terminal.Close()
	if err := unix.SetNonblock(terminal.fd, true); err != nil {
		t.Fatal(err)
	}

	data := make([]byte, 1<<20)
	n, err := terminal.Write(data)
	if n <= 0 || n >= len(data) || err != io.ErrShortWrite {
		t.Fatalf("partial Write = %d, %v; want a partial count and ErrShortWrite", n, err)
	}
	// Keep filling the kernel's output queue until it refuses another write.
	for written := n; written < 4<<20; {
		n, err = terminal.Write(data)
		if n == 0 {
			checkWriteError(t, terminal.name, n, err, unix.EAGAIN)
			return
		}
		if !errors.Is(err, io.ErrShortWrite) {
			t.Fatalf("Write = %d, %v; want ErrShortWrite", n, err)
		}
		written += n
	}
	t.Fatal("output queue did not fill")
}
