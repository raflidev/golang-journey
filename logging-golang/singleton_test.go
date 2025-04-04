package logginggolang

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Helo")
	logrus.Warn("Helo")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Helo")
	logrus.Warn("Helo")

}
