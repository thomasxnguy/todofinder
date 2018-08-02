package todofinder

import (
	"github.com/valyala/fasthttp"
	"github.com/sirupsen/logrus"
	"net"
	. "github.com/thomasxnguy/todofinder/error"
	. "github.com/thomasxnguy/todofinder/http"
	. "github.com/thomasxnguy/todofinder/app"
	"fmt"
)

// Constant defining name of the service.
const (
	ServerName = "todofinder"
)

type Server struct {
	config *Configuration
	log    *logrus.Logger
}

// Init initialize the server for todofinder API
// Configutation and logger are mandatory parameters
func (s *Server) Init(c *Configuration, l *logrus.Logger) error {
	if c == nil || l == nil {
		return fmt.Errorf("failed to initialize server, received nil parameters")
	}
	s.config = c
	s.log = l
	return nil
}

// Run todofinder API server
func (s *Server) Run() error {
	fhttps := &fasthttp.Server{
		Handler: s.handler,
		Name:    ServerName,
	}
	ln, err := s.getListener()
	if err != nil {
		return err
	}

	s.log.WithField("func", "Init").Info("Running Server")
	if s.config.EnableTls {
		if err := fhttps.ServeTLS(ln, s.config.CertFile, s.config.KeyFile); err != nil {
			s.log.WithField("func", "Run").Fatal("Error when serving incoming connections")
			return err
		}
	} else {
		if err := fhttps.Serve(ln); err != nil {
			s.log.WithField("func", "Run").Fatal("Error when serving incoming connections")
			return err
		}
	}
	return nil
}

func (s *Server) getListener() (net.Listener, error) {
	ln, err := net.Listen(s.config.Network, s.config.ListenOn)
	s.log.WithField("func", "Run").Infof("Start server in %s mode and listening to port %s", s.config.Network, s.config.ListenOn)
	if err != nil {
		//log
		return nil, err
	}
	return ln, nil
}

//Run requestHandler with error handling.
func (s *Server) handler(ctx *fasthttp.RequestCtx) {
	defer s.panicHandler(ctx)
	err := Router(ctx)
	if err != nil {
		s.errorHandler(ctx, err)
	}
}

// errorHandler handle the error if request return an error.
// It returns an http response error according to the error code.
func (s *Server) errorHandler(ctx *fasthttp.RequestCtx, e *Error) {
	if e.Error != nil {
		s.log.WithField("func", "errorHandler").Debug("Receive error %v", e)
	}
	if e.ErrorCode == "" {
		e.ErrorCode = INTERNAL_SERVER_ERROR
	}
	msg := e.GetMessage()
	status := Templates[e.ErrorCode].HttpStatus
	errorHttpResponse := &HttpError{e.ErrorCode, msg, status}
	ers, e := errorHttpResponse.ToJson()
	ctx.Error(ers, status)
	//Need to set the content type and headers which have been reset by the previous method
	ctx.SetContentType(CONTENT_TYPE_JSON)
}

// panicHandler prevents the server from exiting from a panic.
// It returns a HTTP 500 error to the end-users.
func (s *Server) panicHandler(ctx *fasthttp.RequestCtx) {
	if rcv := recover(); rcv != nil {
		err := &Error{INTERNAL_SERVER_ERROR, nil, fmt.Errorf("recover from panic: %v", rcv)}
		s.errorHandler(ctx, err)
	}
}
