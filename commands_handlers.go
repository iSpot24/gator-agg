package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/iSpot24/gator-agg/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) > 0 {
		parsed, err := strconv.ParseUint(cmd.args[0], 10, 0)
		if err != nil {
			return err
		}

		if parsed == 0 {
			return errors.New("[<limit>] arg must be positive")
		}
		limit = int32(parsed)
	}
	posts, err := s.db.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  limit,
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		entry := "-> "
		title, err := post.Title.Value()
		if err == nil {
			entry = fmt.Sprintf("%v %v", entry, title)
		}
		url, err := post.Url.Value()
		if err == nil {
			entry = fmt.Sprintf("%v (%v)", entry, url)
		}
		fmt.Println(entry)
	}

	return nil
}

func handlerFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("-> %s (%s)\n", feed.Name, feed.Url)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v", cmd.usage)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("feed created: %v\n", feed.Name)

	err = handlerFollow(s, command{
		name: "follow",
		args: []string{feed.Url},
	}, user)

	if err != nil {
		return err
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v", cmd.usage)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v following %v\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v", cmd.usage)
	}

	err := s.db.DeleteFeedFollowByUserAndFeed(
		context.Background(),
		database.DeleteFeedFollowByUserAndFeedParams{
			ID:  user.ID,
			Url: cmd.args[0],
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("unfollowed successfuly")

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
	}

	return nil
}

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v", cmd.usage)
	}

	freq, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", freq)
	ticker := time.NewTicker(freq)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerUsers(s *state, cmd command, user database.User) error {
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
		return fmt.Errorf("usage: %v", cmd.usage)
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
		return fmt.Errorf("usage: %v", cmd.usage)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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
		return err
	}

	fmt.Println("data deleted")

	return nil
}
