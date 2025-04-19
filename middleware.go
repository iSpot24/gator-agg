package main

import (
	"context"
	"log"

	"github.com/iSpot24/gator-agg/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			log.Fatal("user not logged in")
		}
		return handler(s, cmd, user)
	}
}
