package main

import (
	"fmt"
	"github.com/nibsandpepsi/gator/internal/config"

)
type state struct{
	cfg *config.Config
}

type command struct{
	Name string
	Args []string
}

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) <1 {
		return fmt.Errorf("the login handler expects a single argument, the username.")
	}
	return nil
}