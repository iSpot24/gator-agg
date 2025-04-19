package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
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
	cmds := commands{list: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	return cmds
}
