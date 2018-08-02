package app_test

import (
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

var _ = Describe("Request", func() {
	var (
		ctx *fasthttp.RequestCtx
	)

	BeforeEach(func() {
		ctx = &fasthttp.RequestCtx{}
		s := "../conf/todofinder_test.yaml"
		config, _ := LoadConfiguration(&s)
		log, _ := LoadAppLogger(config)
		ctx.Init(&fasthttp.Request{}, nil, log)
	})

	Describe("Initialize request contextt", func() {
		Context("when using correct parameter", func() {
			It("should not return error", func() {
				error := InitRequest(ctx)
				Expect(error).To(BeNil())
				Expect(GetRequestContext(ctx)).ToNot(BeNil())
			})
		})
	})

})