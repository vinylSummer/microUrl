package httpServer

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(handler http.Handler, options ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	newServer := &Server{
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, option := range options {
		option(newServer)
	}

	newServer.start()

	return newServer
}

func (server *Server) start() {
	server.server.ListenAndServe()

}

func (server *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), server.shutdownTimeout)
	defer cancel()

	return server.server.Shutdown(ctx)
}
