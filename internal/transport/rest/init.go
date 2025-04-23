package rest

import (
	"context"
	"github.com/ysle0/go-starter/internal/db/mongodb"
	"github.com/ysle0/go-starter/internal/server"
)

func InitREST(ctx context.Context, s *server.Server) {
	mongodb.NewMongoDb(ctx, s.Cfg)
	useMiddlewares(s.Router, s.Cfg)
	route(s.Router)
}
