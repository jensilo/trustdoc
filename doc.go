package doc

type Log interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
}
