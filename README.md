# Go Signal Handler

This is a simple library used to place an `os.Signal` value in a channel when it traps a SIGINT or SIGTERM

It's main purpose is for apps that are run within Docker containers to properly handle its signals for shutdown and restart.

## Usage

The library exports one function `Trap()` which returns a `chan os.Signal`

An `os.Signal` value is sent to the channel when either SIGINT or SIGTERM are caught, so that your application may then gracefully shut down.

## Example

```go
import (
	"github.com/zulrang/sighandler"
)

func main() {
	done := sighandler.Trap()
	go func() {
		// your code here
	}
	<- done
	// clean up
}
```

