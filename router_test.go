package todofinder_test

import (
	. "github.com/thomasxnguy/todofinder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/thomasxnguy/todofinder/error"
	. "github.com/thomasxnguy/todofinder/app"
	"github.com/valyala/fasthttp"
)

var _ = Describe("Router", func() {
	var (
		ctx *fasthttp.RequestCtx
	)

	BeforeEach(func() {
		ctx = &fasthttp.RequestCtx{}
		s := "conf/todofinder_test.yaml"
		config, _ := LoadConfiguration(&s)
		log, _ := LoadAppLogger(config)
		ctx.Init(&fasthttp.Request{}, nil, log)
	})

	Describe("Routing request", func() {
		Context("when calling bad request", func() {
			It("should return bad request error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/notexist"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(NOT_FOUND))
			})
		})
	})

	Describe("Routing request", func() {
		Context("when calling not allowed method", func() {
			It("should return method not allowed error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("PUT"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(METHOD_NOT_ALLOWED))
			})
		})
	})

	Describe("Routing request", func() {
		Context("when using bad parameters", func() {
			It("should return bad parameters error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("GET"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(BAD_PARAMETER))
			})
		})
	})

	Describe("Routing request", func() {
		Context("using happy path", func() {
			It("should not return error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("GET"))
				ctx.QueryArgs().Add("package", "fmt")
				ctx.QueryArgs().Add("pattern", "TODO")
				error := Router(ctx)
				Expect(error).To(BeNil())
			})
		})
	})
})
