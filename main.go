package main

import (
	"fmt"
	"os"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/command/handlers"
	"github.com/mgmaster24/gator/internal/config"
)

func main() {
	gatorConfig := config.Read()
	var state internal.State = internal.State{
		Cfg: &gatorConfig,
	}

	var commands *command.Commands = &command.Commands{
		CmdMap: make(map[string]func(*internal.State, command.Command) error),
	}

	commands.Register("login", handlers.Login)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	cmd := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	err := commands.Run(&state, cmd)
	if err != nil {
		fmt.Println("Failed to run command. Error", err)
		os.Exit(1)
	}
}
