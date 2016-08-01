package sighandler

import (
	"testing"
	"time"
	"os"
	"syscall"
)

func TestTrap(t *testing.T) {
	done := Trap()
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf(err.Error())
	}

	// test sigint trapping
	p.Signal(syscall.SIGINT)
	select {
	case _ = <-done:
	case <-time.After(time.Second * 2):
		t.Fatalf("Done signal not received from SIGINT")
	}

	// test sigterm trapping
	p.Signal(syscall.SIGTERM)
	select {
	case _ = <-done:
	case <-time.After(time.Second * 2):
		t.Fatalf("Done signal not received from SIGTERM")
	}
}