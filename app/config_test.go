package app_test

import (
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		configPath = "../conf/todofinder_test.yaml"
	)

	Describe("Validate configuration file", func() {
		Context("with enable TLS without certificate", func() {
			It("should return error", func() {
				config := &Configuration{}
				config.EnableTls = true
				Expect(config.Validate()).NotTo(BeNil())
			})
		})
	})

	Describe("Validate configuration file", func() {
		Context("with log to file without location specified", func() {
			It("should return error", func() {
				config := &Configuration{}
				config.LogOutput = "file"
				Expect(config.Validate()).NotTo(BeNil())
			})
		})
	})

	Describe("Load configuration file", func() {
		Context("when file exist", func() {
			It("should not return error", func() {
				config, err := LoadConfiguration(&configPath)
				Expect(err).To(BeNil())
				Expect(config.ListenOn).To(Equal(":8089"))
				Expect(config.Network).To(Equal("tcp4"))
				Expect(config.EnableTls).To(Equal(false))
			})
		})
	})

	Describe("Load configuration file", func() {
		Context("when file does not exit", func() {
			It("should return error", func() {
				configPath = "../conf/3adea5d7-19ec-4fff-869d-8aff8b513311"
				_, err := LoadConfiguration(&configPath)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
