package app

import (
	"github.com/valyala/fasthttp"
)

// InitRequest initialize the request with a request context and log access.
func InitRequest(ctx *fasthttp.RequestCtx) error {
	ac := NewRequestContext(ctx)
	ctx.SetUserValue("Context", ac)
	logAccess(ctx)
	return nil

}

// GetRequestScope returns the RequestContext of the current request.
func GetRequestContext(ctx *fasthttp.RequestCtx) RequestContext {
	return ctx.UserValue("Context").(RequestContext)
}

// logAccess logs a message describing the current request.
func logAccess(ctx *fasthttp.RequestCtx) {
	GetRequestContext(ctx).Infof("starting processing request")
}
