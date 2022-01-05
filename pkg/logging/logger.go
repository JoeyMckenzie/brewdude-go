package logging

type Logger interface {
	Info(message string)
	Infof(message string, context ...string)
	ErrorWhile(message string, error error)
	Errorf(message string, context ...string)
	FatalWhile(message string, error error)
}
