package term

import (
	"log"
	"time"
)

// Reset an Arduino by toggling the DTR signal.
func ExampleStatus_SetDTR() {
	t, _ := Open("/dev/USB0")

	status, _ := t.Status()
	status.SetDTR(!status.DTR())
	t.SetStatus(status)

	time.Sleep(1 * time.Second)

	status.SetDTR(!status.DTR())
	t.SetStatus(status)
}

// Send Break to the remote DTE.
func ExampleTerm_SendBreak() {
	t, _ := Open("/dev/ttyUSB1")
	for {
		time.Sleep(3 * time.Second)
		log.Println("Break...")
		t.SendBreak()
	}
}
