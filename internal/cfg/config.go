package cfg

import (
	"errors"
	"fmt"
	"github.com/naoina/toml"
	"log"
	"os"
)

var (
	InvalidConfigPathErr = errors.New("invalid cfg path")
)

type Config struct {
	Server struct {
		Port                       string
		Info                       string
		HandlerResponseTimeoutInMs int
	}
}

func NewConfig(path string) (Config, error) {
	cfg := Config{}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, InvalidConfigPathErr
	}
	defer file.Close()

	dec := toml.NewDecoder(file)
	err = dec.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) DbgPrint() {
	s := fmt.Sprintf("used config -> %#v\n", c)
	log.Print(s)
}
