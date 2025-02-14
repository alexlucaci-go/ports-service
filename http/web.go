package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type Service struct {
	shutdown chan os.Signal
	mux      *http.ServeMux
}

func NewService(shutdown chan os.Signal) *Service {
	return &Service{
		shutdown: shutdown,
		mux:      http.NewServeMux(),
	}
}

func (s *Service) Handle(method string, path string, handler Handler) {

	handler = errorsMid()(handler)
	// add other middlewares here
	// handler = loggerMid()(handler)
	// ...

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			log.Printf("Uncatched error : %v\n", err)
			return
		}
	}

	// since go 1.22 http.HandleFunc also supports the verb
	pathWithMethod := fmt.Sprintf("%s %s", method, path)
	s.mux.HandleFunc(pathWithMethod, h)
}

func (s *Service) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Service) SignalShutdown() {
	s.shutdown <- syscall.SIGTERM
}
