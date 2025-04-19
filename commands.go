package main

import (
	"fmt"
)

type command struct {
	name  string
	args  []string
	usage string
}

type commands struct {
	list   map[string]func(*state, command) error
	usages map[string]string
}

func (c *commands) register(name string, usage string, f func(s *state, cmd command) error) {
	c.list[name] = f
	c.usages[name] = usage
}

func (c *commands) run(s *state, cmd command) error {
	wantedCmd, exists := c.list[cmd.name]
	if !exists {
		return fmt.Errorf("command %s not found", cmd.name)
	}

	err := wantedCmd(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func initCommands() commands {
	cmds := commands{
		list:   make(map[string]func(*state, command) error),
		usages: make(map[string]string),
	}

	cmds.register("login", "login <username>", handlerLogin)
	cmds.register("register", "register <username>", handlerRegister)
	cmds.register("reset", "reset", handlerReset)
	cmds.register("users", "users", middlewareLoggedIn(handlerUsers))
	cmds.register("agg", "agg <time_freq_string>", middlewareLoggedIn(handlerAgg))
	cmds.register("addfeed", "addfeed <feed_name> <feed_url>", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", "feeds", middlewareLoggedIn(handlerFeeds))
	cmds.register("follow", "follow <feed_url>", middlewareLoggedIn(handlerFollow))
	cmds.register("unfollow", "unfollow <feed_url>", middlewareLoggedIn(handlerUnfollow))
	cmds.register("following", "following", middlewareLoggedIn(handlerFollowing))
	cmds.register("browse", "browse <limit>", middlewareLoggedIn(handlerBrowse))

	return cmds
}
