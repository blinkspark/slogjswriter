# slogjswriter

中文 / Chinese: 本仓库的主要 README 为中文；若你需要英文说明，请参见 `README_en.md`。

一个将写入数据同时发布到 NATS JetStream 的简单 `io.Writer` 实现，方便把日志或其他流式数据同步到 JetStream。

安装

```bash
go get github.com/blinkspark/slogjswriter
```

快速示例

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
    // 将 w 传给日志库，或直接写入
    _, _ = w.Write([]byte("hello jetstream\n"))
}
```

TODO

- 添加单元测试（可以用 test 目录中的模拟对象替代真实 JetStream 连接）。

License: MIT