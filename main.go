package sighandler

import (
	"log"
    "os"
	"os/signal"
	"syscall"
)

// Trap creates traps for SIGINT and SIGTERM and returns a channel which 
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
		sig := <-sigs
		// log signal type
		log.Println(sig)
		done <- true
	}()

	// return channel to listen on
	return done
}
