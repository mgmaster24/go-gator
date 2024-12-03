package handlers

import (
	"context"
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
)

func Users(s *internal.State, cmd command.Command) error {
	users, err := s.Queries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error get the list of users Error: %s", err.Error())
	}

	for _, user := range users {
		name := user
		if user == s.Cfg.CurrentUserName {
			name += " (current)"
		}

		fmt.Println(name)
	}
	return nil
}
