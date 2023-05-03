package util

type Logger interface {
	Debug(message string)
	DebugWithCtx(message string, ctx map[string]any)

	Info(message string)
	InfoWithCtx(message string, ctx map[string]any)

	Warn(message string)
	WarnWithCtx(message string, ctx map[string]any)

	Error(message string)
	ErrorWithCtx(message string, ctx map[string]any)
}
