package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ip       string
	port     int
	listener net.Listener
}

func NewServer(port int) (*Server, error) {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("tcp listen error: %v", err)
	}
	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     port,
		listener: listener,
	}, nil
}

func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	go func() {
		slog.Info("v1 server listening at", "addr", s.ip+strconv.Itoa(s.port))
		if err := srv.Serve(s.listener); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "err", err)
		}
	}()

	<-ctx.Done()
	slog.Info("shutdown signal recieved")

	// Give requests 30 seconds to complete
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutDownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}

	return nil
}

func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,

		// For request headers
		ReadHeaderTimeout: 3 * time.Second,
		// Keep-alive connection time
		IdleTimeout: 120 * time.Second,
	})
}
