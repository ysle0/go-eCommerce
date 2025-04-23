package rest

import (
	"context"
	"github.com/ysle0/go-starter/internal/db/mongo"
	"github.com/ysle0/go-starter/internal/server"
)

func InitREST(ctx context.Context, s *server.Server) {
	mongo.NewMongo(ctx, s.Cfg)
	useMiddlewares(s.Router, s.Cfg)
	route(s.Router)
}
