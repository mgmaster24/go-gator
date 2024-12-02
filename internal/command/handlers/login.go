package handlers

import (
	"context"
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
)

func Login(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expects a single parameter for the user name. command: %s", cmd)
	}

	user, err := s.Queries.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error retrieving user %s. Error: %s", cmd.Args[0], err.Error())
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has logged in.\n", user.Name)
	return nil
}
