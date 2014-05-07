// Package termios provides a low level interface to the termios(3)
// terminal line discipline facilities.
//
// For a higher level interface please use the github.com/pkg/term package.
package termios

import (
	"syscall"
)

const (
	// The change occurs immediately.
	TCSANOW = syscall.TCSANOW

	// The change occurs after all output written to fildes has been
	// transmitted to the terminal.  This value of optional_actions
	// should be used when changing parameters that affect output.
	TCSADRAIN = syscall.TCSADRAIN

	//  The change occurs after all output written to fildes has been
	//         transmitted to the terminal.  Additionally, any input that has
	//         been received but not read is discarded.
	TCSAFLUSH = syscall.TCSAFLUSH

	// If this value is or'ed into the optional_actions value, the
	//        values of the c_cflag, c_ispeed, and c_ospeed fields are
	//       ignored.
	TCSASOFT = syscall.TCSASOFT
)
