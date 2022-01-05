package logging

import (
	"log"
	"os"
)

type consoleLogger struct {
	logger *log.Logger
}

func New(prefix string) Logger {
	return consoleLogger{
		log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

func (cl consoleLogger) Info(message string) {
	cl.logger.Printf("INFO: %s", message)
}

func (cl consoleLogger) Infof(message string, context ...string) {

}

func (cl consoleLogger) ErrorWhile(message string, error error) {
	cl.logger.Printf("ERROR: An occurred while %s - %s", message, error.Error())
}

func (cl consoleLogger) Errorf(message string, context ...string) {

}

func (cl consoleLogger) FatalWhile(message string, error error) {
	cl.logger.Fatalf("An unrecoverable error while %s - %s", message, error.Error())
}
