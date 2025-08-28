package slogjswriter

import (
	"context"
	"io"

	"github.com/nats-io/nats.go/jetstream"
)

// JetStreamWriter implements io.Writer and mirrors written bytes to a JetStream
// subject while optionally writing them to a local io.Writer.
//
// It's suitable for plugging into logging frameworks that accept an io.Writer.
//
// JetStreamWriter 实现了 io.Writer，会将写入的字节同时转发到 JetStream
// 的 subject，并可选地写入本地的 io.Writer。
//
// 适用于需要传入 io.Writer 的日志框架。
type JetStreamWriter struct {
	writer  io.Writer
	stream  jetstream.JetStream
	subject string
}

var _ io.Writer = &JetStreamWriter{}

// NewJetStreamWriter creates a JetStreamWriter that publishes
// to the provided JetStream subject.
//
// NewJetStreamWriter 创建一个 JetStreamWriter，将log发布到
// 指定的 JetStream subject。
func NewJetStreamWriter(stream jetstream.JetStream, subject string) (*JetStreamWriter, error) {
	return NewJetStreamWriterWithWriter(stream, subject, nil)
}

// NewJetStreamWriterWithWriter creates a JetStreamWriter using the provided
// io.Writer as the local writer (useful for testing or redirecting output).
//
// NewJetStreamWriterWithWriter 使用提供的 io.Writer 作为本地写入器创建
// JetStreamWriter（在测试或重定向输出时有用）。
func NewJetStreamWriterWithWriter(stream jetstream.JetStream, subject string, w io.Writer) (*JetStreamWriter, error) {
	return &JetStreamWriter{writer: w, stream: stream, subject: subject}, nil
}

// Write writes p to the configured local writer and publishes p to JetStream.
// It returns the number of bytes written to the local writer and any error
// encountered (either from the local write or from the publish operation).
//
// Write 将 p 写入配置的本地写入器，并将 p 发布到 JetStream。
// 它返回写入到本地写入器的字节数以及遇到的任何错误（来自本地写入或发布操作）。
func (s *JetStreamWriter) Write(p []byte) (n int, err error) {
	if s.writer != nil {
		n, err = s.writer.Write(p)
		if err != nil {
			return n, err
		}
	}

	// Try to publish to JetStream. If publish fails, return the publish error
	// so callers are aware the remote publish didn't succeed.
	//
	// 尝试发布到 JetStream。如果发布失败，返回发布错误，
	// 以便调用方知道远程发布未成功。
	_, err = s.stream.Publish(context.Background(), s.subject, p)
	return len(p), err
}
