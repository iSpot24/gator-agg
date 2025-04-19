package main

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/iSpot24/gator-agg/internal/config"
	"github.com/iSpot24/gator-agg/internal/database"
	"github.com/iSpot24/gator-agg/internal/feeder"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	cfg    *config.Config
	client feeder.Client
}

func listenCmd(cmds *commands, state *state) error {
	args := os.Args
	if len(args) < 2 {
		return errors.New("not enough arguments provided")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	db, err := sql.Open("postgres", state.cfg.DbURL)

	if err != nil {
		return err
	}
	defer db.Close()

	state.db = database.New(db)
	state.client = feeder.NewClient(20 * time.Second)

	return cmds.run(state, cmd)
}
