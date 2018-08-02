package app

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/http"
)

// RequestContext contains the application-specific information that are carried around in a request.
type RequestContext interface {
	Logger
	RemoteAddress() string
	RequestID() string
	Now() time.Time
}

type requestContext struct {
	Logger                  // the logger tagged with the current request information
	remoteAddress string    // the remote address of the request
	requestID     string    // an unique ID identifying one or multiple correlated HTTP requests
	now           time.Time // the time when the request is being processed
}

// RemoteAddress returns the IP address of the requester.
func (rs *requestContext) RemoteAddress() string {
	return rs.remoteAddress
}

// Now returns time when the request is being processed, used for concurrency operations.
func (rs *requestContext) Now() time.Time {
	return rs.now
}

// RequestID returns the ID of the current request.
func (rs *requestContext) RequestID() string {
	return rs.requestID
}

// NewRequestContext creates a new RequestContext with the current request information.
func NewRequestContext(ctx *fasthttp.RequestCtx) RequestContext {
	now := time.Now()
	requestUri := string(ctx.URI().RequestURI())
	remoteIp := ctx.RemoteIP().String()
	requestID := string(ctx.Request.Header.Peek(X_REQUEST_ID))

	l := NewLogger(logrus.Fields{})
	l.SetField("requestUri", requestUri)
	l.SetField("remoteIp", remoteIp)
	l.SetField("userAgent", string(ctx.UserAgent()))
	if requestID != "" {
		l.SetField(X_REQUEST_ID, requestID)
	}
	return &requestContext{
		Logger:        l,
		remoteAddress: remoteIp,
		now:           now,
		requestID:     requestID,
	}
}
