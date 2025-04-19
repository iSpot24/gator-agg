package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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

	if usage, exists := cmds.usages[cmd.name]; exists {
		cmd.usage = usage
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

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	res, err := s.client.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	for _, item := range res.Channel.Item {
		postTimestamp, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			log.Fatal(err)
		}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       sql.NullString{String: item.Title, Valid: true},
			Description: sql.NullString{String: item.Description, Valid: true},
			Url:         sql.NullString{String: item.Link, Valid: true},
			FeedID:      feed.ID,
			PublishedAt: sql.NullTime{Time: postTimestamp, Valid: true},
		})

		if err != nil {
			log.Fatal(err)
		}

		if title, err := post.Title.Value(); err == nil {
			fmt.Printf("Saved: %v\n", title)
		}
	}

	return nil
}
