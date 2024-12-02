package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
)

func Register(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("register expects a single parameter for the user. command: %s", cmd)
	}

	uuid := uuid.New()
	now := time.Now()
	name := cmd.Args[0]

	user, err := s.Queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been registered.\n", user.Name)
	return nil
}
