package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger) Debug(message string) {
	log.Println("[debug] " + message)
}

func (l Logger) Info(message string) {
	log.Println("[info] " + message)
}

func (l Logger) Warn(message string) {
	log.Println("[warn] " + message)
}

func (l Logger) Error(message string) {
	log.Println("[error] " + message)
}

func (l Logger) DebugWithCtx(message string, ctx map[string]any) {
	log.Printf("[debug] %s %s", message, ctxToString(ctx))
}

func (l Logger) InfoWithCtx(message string, ctx map[string]any) {
	log.Printf("[info] %s %s", message, ctxToString(ctx))
}

func (l Logger) WarnWithCtx(message string, ctx map[string]any) {
	log.Printf("[warn] %s %s", message, ctxToString(ctx))
}

func (l Logger) ErrorWithCtx(message string, ctx map[string]any) {
	log.Printf("[error] %s %s", message, ctxToString(ctx))
}

func ctxToString(ctx map[string]any) string {
	var buf bytes.Buffer

	for k, v := range ctx {
		buf.Write([]byte(fmt.Sprintf("%v=%v ", k, v)))
	}

	s, _ := io.ReadAll(&buf)
	return string(s)
}
