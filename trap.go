// Package sighandler is a SIGTERM and SIGINT handler for Go programs
package sighandler

import (
    "os"
	"os/signal"
	"syscall"
)

// Trap handles SIGINT and SIGTERM signals and returns a channel which 
// will produce true when one of those signals are trapped.
func Trap() chan bool {
	// signal channel
	sigs := make(chan os.Signal, 1)
	// listen channel
	done := make(chan bool)

	// trap signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// signal handler
	go func() {
		// allow signals to be continually trapped
		for {
			_ = <-sigs
			// log signal type
			done <- true
		}
	}()

	// return channel to listen on
	return done
}
