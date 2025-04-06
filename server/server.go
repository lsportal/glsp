package server

import (
	"time"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
)

type CurrentConnections map[string]*jsonrpc2.Conn

var (
	DefaultTimeout     = time.Minute
	currentConnections = make(CurrentConnections)
)

//
// Server
//

type Server struct {
	Handler            glsp.Handler
	LogBaseName        string
	Debug              bool
	CurrentConnections map[string]*jsonrpc2.Conn

	Log              commonlog.Logger
	Timeout          time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	StreamTimeout    time.Duration
	WebSocketTimeout time.Duration
}

func NewServer(handler glsp.Handler, logName string, debug bool) *Server {
	return &Server{
		Handler:            handler,
		LogBaseName:        logName,
		Debug:              debug,
		CurrentConnections: currentConnections,
		Log:                commonlog.GetLogger(logName),
		Timeout:            DefaultTimeout,
		ReadTimeout:        DefaultTimeout,
		WriteTimeout:       DefaultTimeout,
		StreamTimeout:      DefaultTimeout,
		WebSocketTimeout:   DefaultTimeout,
	}
}
