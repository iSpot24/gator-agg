package main

import (
	"log"

	"github.com/iSpot24/gator-agg/internal/config"
)

func main() {
	var cfg config.Config
	err := cfg.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cfgState := state{cfg: &cfg}
	cmds := initCommands()

	err = listenCmd(&cmds, &cfgState)

	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
