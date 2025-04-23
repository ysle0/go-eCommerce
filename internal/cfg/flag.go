package cfg

import (
	"flag"
	"log"
	"os"
)

func GetConfigPath() (string, error) {
	if len(os.Args) == 0 {
		log.Fatal("web need arguments")
	}

	var cfgPath string
	fs := flag.NewFlagSet("web", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	fs.StringVar(&cfgPath, "cfg", "", "cfg path")

	s := os.Args[1:]
	err := fs.Parse(s)

	return cfgPath, err
}
