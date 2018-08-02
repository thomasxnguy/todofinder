package app_test

import (
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/http"
	. "github.com/onsi/ginkgo"
	"github.com/valyala/fasthttp"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context", func() {
	var (
		fastHttpCtx *fasthttp.RequestCtx
	)

	BeforeEach(func() {
		fastHttpCtx = &fasthttp.RequestCtx{}
		s := "conf/todofinder_test.yaml"
		config, _ := LoadConfiguration(&s)
		log, _ := LoadAppLogger(config)
		fastHttpCtx.Init(&fasthttp.Request{}, nil, log)
		fastHttpCtx.Request.SetRequestURIBytes([]byte("/test"))
		fastHttpCtx.Request.Header.Add(X_REQUEST_ID, "693860d3-8b9c-42bc-aab2-f26c96533292")

	})

	Describe("Creating new Request", func() {
		Context("with valid call", func() {
			It("should return context", func() {
				ctx := NewRequestContext(fastHttpCtx)
				Expect(ctx.RequestID()).To(Equal("693860d3-8b9c-42bc-aab2-f26c96533292"))
				Expect(ctx.RemoteAddress()).To(Equal("0.0.0.0"))
				Expect(ctx.Now()).ToNot(BeNil())
			})
		})
	})
})
