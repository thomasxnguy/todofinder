package todofinder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

// TestApp is the main function to run the test suite for "todofinder" package following gingko testing framework
func TestTodofinder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Todofinder Suite")
}
