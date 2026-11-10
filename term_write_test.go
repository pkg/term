// +build !windows

package term

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func TestTermWrite(t *testing.T) {
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

	for _, data := range [][]byte{nil, {}, []byte("terminal output\n")} {
		n, err := terminal.Write(data)
		if n != len(data) || err != nil {
			t.Fatalf("Write(%q) = %d, %v", data, n, err)
		}
		if len(data) > 0 {
			got := make([]byte, len(data))
			if _, err := io.ReadFull(master, got); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, data) {
				t.Fatalf("output = %q, want %q", got, data)
			}
		}
	}

	if err := terminal.Close(); err != nil {
		t.Fatal(err)
	}
	for _, data := range [][]byte{nil, {}, []byte("closed terminal")} {
		n, err := terminal.Write(data)
		checkWriteError(t, terminal.name, n, err, unix.EBADF)
	}
}

func checkWriteError(t *testing.T, name string, n int, err error, want error) {
	t.Helper()
	var pathErr *os.PathError
	if n != 0 || !errors.Is(err, want) || !errors.As(err, &pathErr) {
		t.Fatalf("Write = %d, %v; want 0 and PathError wrapping %v", n, err, want)
	}
	if pathErr.Op != "write" || pathErr.Path != name {
		t.Fatalf("PathError = %#v", pathErr)
	}
}
