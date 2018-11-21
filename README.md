# Style Logger

Add styles to Go standard `log.Logger`.

# Installation

```sh
go get github.com/haobaozhong/stylelog
```

# Example

```go
package main

import (
	"log"
	"os"

	"github.com/haobaozhong/stylelog"
)

func main() {
	baseLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	styleLogger := stylelog.New(baseLogger)

	styleLogger.SetStyle(stylelog.Bold, stylelog.Underline, stylelog.Green, stylelog.Faint)
	styleLogger.Print("hello, world")
	styleLogger.Reset()
}
```

## Feedback

Feedback is greatly appreciated.
Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing

Contributions are greatly appreciated.
The maintainers actively manage the issues list, and try to highlight issues suitable for newcomers.
The project follows the typical GitHub pull request model.
