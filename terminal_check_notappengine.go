// +build !appengine,!js,!nacl,!plan9,!windows

package logrus

import (
	"io"
	"os"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}
