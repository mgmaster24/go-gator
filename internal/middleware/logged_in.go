package middleware

import (
	"context"
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
)

func LoggedIn(
	handler func(s *internal.State, cmd command.Command, user database.User) error,
) func(*internal.State, command.Command) error {
	return func(s *internal.State, c command.Command) error {
		user, err := s.Queries.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			fmt.Println("Error getting the current user")
			return err
		}

		return handler(s, c, user)
	}
}
