package db

import (
	"github.com/ysle0/go-starter/internal/cfg"
)

type Repo struct {
	cfg cfg.Config
}

func NewRepository(cfg cfg.Config) *Repo {
	r := &Repo{
		cfg: cfg,
	}

	return r
}
