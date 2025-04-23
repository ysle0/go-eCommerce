package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/ysle0/go-starter/internal/cfg"
	"log"
	"net/http"
)

type Server struct {
	Cfg    cfg.Config
	Router *chi.Mux
}

func NewServer(cfg cfg.Config, r *chi.Mux) *Server {
	s := &Server{
		Cfg:    cfg,
		Router: r,
	}

	return s
}

func (s *Server) Serve() error {
	port := s.Cfg.Server.Port
	log.Printf("serve at %s", port)
	return http.ListenAndServe(port, s.Router)
}
