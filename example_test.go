package term

import (
	"log"
	"time"
)

// Reset an Arduino by lowering the DTR signal.
func ExampleStatusSetDTR(t *Term) {
	status, err := t.Status()
	if err != nil {
		log.Fatal(err)
	}
	status.SetDTR(false)
	if err := t.SetStatus(status); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	status.SetDTR(true)
	if err := t.SetStatus(status); err != nil {
		log.Fatal(err)
	}
}
