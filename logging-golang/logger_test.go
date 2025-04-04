package logginggolang

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello logger")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("applicaition.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.SetOutput(file)

	logger.Info("Hello logger")
	logger.Warn("Hello logger")
	logger.Error("Hello logger")

}

func TestFormat(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logger")
	logger.Warn("Hello logger")
	logger.Error("Hello logger")
}
