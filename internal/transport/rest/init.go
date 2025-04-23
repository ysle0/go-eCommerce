package rest

import "github.com/ysle0/go-starter/internal/server"

func InitREST(s *server.Server) {
	useMiddlewares(s.Router, s.Cfg)
	route(s.Router)
}
