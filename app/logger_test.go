package app_test

import (
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var _ = Describe("Logger", func() {
	var (
		config *Configuration
	)

	BeforeEach(func() {
		s := "../conf/todofinder_test.yaml"
		config, _ = LoadConfiguration(&s)

	})

	Describe("Loading logger", func() {
		Context("with valid configuration file", func() {
			It("should return logger", func() {
				log, err := LoadAppLogger(config)
				Expect(err).To(BeNil())
				Expect(log.Level).To(Equal(logrus.InfoLevel))
				Expect(log.Out).To(Equal(ioutil.Discard))
			})
		})
	})
})
