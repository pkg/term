package term

import (
	"log"
	"time"
)

// Reset an Arduino by toggling the DTR signal.
func ExampleTerm_SetDTR() {
	t, _ := Open("/dev/USB0")

	t.SetDTR(false)

	time.Sleep(1 * time.Second)

	t.SetDTR(true)
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
