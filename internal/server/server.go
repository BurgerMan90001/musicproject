package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
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
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		shutDownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		errCh <- srv.Shutdown(shutDownCtx)
	}()

	log.Printf("Server listening at %s. v1\n", s.ip)
	if err := srv.Serve(s.listener); err != nil {
		return err
	}

	if err := <-errCh; err != nil {
		return err
	}

	return nil
}
func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,
	})
}

