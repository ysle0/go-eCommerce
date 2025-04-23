package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/ysle0/go-starter/internal/cfg"
	"github.com/ysle0/go-starter/internal/server"
	"github.com/ysle0/go-starter/internal/transport/rest"
	"log"
)

func main() {
	cfgPath, err := cfg.GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	c, err := cfg.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	//c.DbgPrint()

	ctx := context.Background()
	s := server.NewServer(c, chi.NewRouter())
	rest.InitREST(ctx, s)
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
