package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
	"github.com/mgmaster24/gator/internal/rss"
)

func Aggregate(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf(
			"Aggregate requires the an argument for time between requests. Command: %v",
			cmd,
		)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collectiong feeds every %f seconds\n", timeBetweenRequests.Seconds())
	for ; ; <-ticker.C {
		feed, err := s.Queries.GetNextFeedToFetch(context.Background())
		if err != nil {
			return err
		}

		_, err = s.Queries.MarkFeedFetched(context.Background(), feed.ID)
		if err != nil {
			return err
		}

		fmt.Println("-------------------------- Fetching Feed ----------------------------")
		rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
		if err != nil {
			return err
		}

		fmt.Println("\tFeed Title:", rssFeed.Channel.Title)
		fmt.Println("\tFeed URL:", rssFeed.Channel.Link)

		for _, item := range rssFeed.Channel.Item {
			publishedAt := sql.NullTime{}
			if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
				publishedAt = sql.NullTime{
					Time:  t,
					Valid: true,
				}
			}
			s.Queries.CreatePost(context.Background(), database.CreatePostParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title:     item.Title,
				Url: sql.NullString{
					String: item.Link,
					Valid:  true,
				},
				Description: sql.NullString{
					String: item.Description,
					Valid:  true,
				},
				FeedID:      feed.ID,
				PublishedAt: publishedAt,
			})
		}

		fmt.Printf("\tFeed %s collected. Found %v posts\n", feed.Name, len(rssFeed.Channel.Item))
		fmt.Println("--------------------------- Fetched Feed -----------------------------")
	}
}
