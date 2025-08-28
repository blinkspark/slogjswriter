# slogjswriter
若你需要中文说明，请参见 `README_zhCN.md`。

English: This repository provides a simple `io.Writer` implementation that publishes written data to NATS JetStream, making it easy to forward logs or other streaming data to JetStream.

Installation

```bash
go get github.com/blinkspark/slogjswriter
```

Quick example

```go
package main

import (
    "os"
    "github.com/blinkspark/slogjswriter"
    "github.com/nats-io/nats.go"
)

func main() {
    nc, _ := nats.Connect(nats.DefaultURL)
    js, _ := nc.JetStream()
    w, _ := slogjswriter.NewJetStreamWriterWithWriter(js, "logs.subject", os.Stdout)
    // Pass w to a logging library, or write directly
    _, _ = w.Write([]byte("hello jetstream\n"))
}
```

TODO

- Add unit tests (use test doubles to mock JetStream where appropriate).

License: MIT
