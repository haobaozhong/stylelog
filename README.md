# Style Logger

style logger add styles to Go standard `log.Logger`.

# Installation

```sh
go get github.com/haobaozhong/stylelog
```

# Example

```go
baseLogger := log.New(os.Stdout, "STYLE LOGGER: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
styleLogger := stylelog.New(baseLogger)

logger.Style(Bold, Underline, Green, Faint)
logger.Print("hello, world")
logger.Reset()
```

## Feedback

Feedback is greatly appreciated.
Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing

Contributions are greatly appreciated.
The maintainers actively manage the issues list, and try to highlight issues suitable for newcomers.
The project follows the typical GitHub pull request model.
