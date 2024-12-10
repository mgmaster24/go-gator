package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
)

func Browse(s *internal.State, cmd command.Command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if val, err := strconv.Atoi(cmd.Args[0]); err != nil {
			fmt.Println("Error parsing command value. Using default limit of 2")
		} else {
			limit = val
		}
	}

	posts, err := s.Queries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url.String)
		fmt.Println("=====================================")
	}

	return nil
}
