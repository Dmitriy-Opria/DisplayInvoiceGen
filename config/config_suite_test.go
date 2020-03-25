package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"testing"
)

func TestServices(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
