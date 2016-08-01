// Package sighandler is a SIGTERM and SIGINT handler for Go programs
package sighandler

import (
    "os"
	"os/signal"
	"syscall"
)

// Trap handles SIGINT and SIGTERM signals and returns a channel which 
// will produce the os.Signal when one of those signals are trapped.
func Trap() chan os.Signal {
	// signal channel
	sigs := make(chan os.Signal, 1)
	// listen channel
	done := make(chan os.Signal)

	// trap signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// signal handler
	go func() {
		// allow signals to be continually trapped
		for {
			done <- <-sigs
		}
	}()

	// return channel to listen on
	return done
}
