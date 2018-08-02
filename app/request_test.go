package app_test

import (
	. "github.com/thomasxnguy/todofinder/app"

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