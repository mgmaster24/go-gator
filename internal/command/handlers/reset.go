package handlers

import (
	"context"
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
)

func Reset(s *internal.State, cmd command.Command) error {
	_, err := s.Queries.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting the users table Error: %s", err.Error())
	}

	fmt.Println("Reset all users in the users table")
	return nil
}
