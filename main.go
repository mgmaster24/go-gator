package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/command/handlers"
	"github.com/mgmaster24/gator/internal/config"
	"github.com/mgmaster24/gator/internal/database"
	"github.com/mgmaster24/gator/internal/middleware"
)

func main() {
	gatorConfig := config.Read()
	db, err := sql.Open("postgres", gatorConfig.DbUrl)
	var state internal.State = internal.State{
		Cfg:     &gatorConfig,
		Queries: database.New(db),
	}

	var commands *command.Commands = &command.Commands{
		CmdMap: make(map[string]func(*internal.State, command.Command) error),
	}

	commands.Register("login", handlers.Login)
	commands.Register("register", handlers.Register)
	commands.Register("reset", handlers.ResetUsers)
	commands.Register("users", handlers.Users)
	commands.Register("agg", handlers.Aggregate)
	commands.Register("addfeed", middleware.LoggedIn(handlers.AddFeed))
	commands.Register("feeds", handlers.Feeds)
	commands.Register("follow", middleware.LoggedIn(handlers.Follow))
	commands.Register("unfollow", middleware.LoggedIn(handlers.Unfollow))
	commands.Register("following", middleware.LoggedIn(handlers.Following))
	commands.Register("browse", middleware.LoggedIn(handlers.Browse))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	cmd := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Printf("Failed to run command %s.\nError: %s\n", cmd.Name, err.Error())
		os.Exit(1)
	}
}
