package todofinder_test

import (
	. "github.com/thomasxnguy/todofinder"
	. "github.com/thomasxnguy/todofinder/error"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Search", func() {
	var (
		packageName = "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder"
		pattern     = "TODO"
	)

	Describe("Import package", func() {
		Context("on an existing package", func() {
			It("should not return error", func() {
				dir, _ := os.Getwd()
				_, error := ImportPkg(packageName, dir)
				Expect(error).To(BeNil())
			})
		})
	})

	Describe("Import package", func() {
		Context("on an non existing package", func() {
			It("should return package not found error", func() {
				dir, _ := os.Getwd()
				_, error := ImportPkg("0d37c866-d6ec-4a8a-85a2-29ced9683d2f", dir)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(PACKAGE_NOT_FOUND))
			})
		})
	})

	Describe("Import package", func() {
		Context("by calling a command package", func() {
			It("should return no source error", func() {
				dir, _ := os.Getwd()
				_, error := ImportPkg("./cmd", dir)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(NO_SOURCE))
			})
		})
	})

	Describe("Extract pattern", func() {
		Context("when pattern exist", func() {
			It("should result a result", func() {
				dir, _ := os.Getwd()
				p, _ := ImportPkg(packageName, dir)
				rch := make(chan *SearchResult, 10)
				go ExtractPattern(p, pattern, rch)
				searchResult := <-rch
				Expect(searchResult).NotTo(BeNil())
			})
		})
	})
})
