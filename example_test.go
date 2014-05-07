package term

import (
	"log"
	"time"
)

// Reset an Arduino by toggling the DTR signal.
func ExampleStatus_SetDTR() {
	t, _ := Open("/dev/USB0")
	status, _ := t.Status()
	status.SetDTR(false)

	t.SetStatus(status)
	time.Sleep(1 * time.Second)

	status.SetDTR(true)
	t.SetStatus(status)
}
