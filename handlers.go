package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iSpot24/gator-agg/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := s.client.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New("no users found")
	}

	fmt.Printf("%+v\n", rssFeed)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New("no users found")
	}

	for _, user := range users {
		name := user.Name
		if s.cfg.Username == name {
			name += " (current)"
		}

		fmt.Printf("* %s\n", name)
	}

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("user not found")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("logged in as %s\n", user.Name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return errors.New("user already registered")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("user registered")

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())

	if err != nil {
		return errors.New("users table could not be deleted")
	}

	fmt.Println("users table data deleted")

	return nil
}
