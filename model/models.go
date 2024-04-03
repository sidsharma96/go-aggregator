package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/sidsharma96/go-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type UserIdsAndKeys struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

func DatabaseUsersToIds(userIdsAndKeys []database.GetAllUserIdsAndKeysRow) []UserIdsAndKeys {
	result := make([]UserIdsAndKeys, len(userIdsAndKeys))
	for i, userIdAndKey := range userIdsAndKeys {
		result[i] = UserIdsAndKeys {
			Name: userIdAndKey.Name,
			ID: userIdAndKey.ID,
			ApiKey: userIdAndKey.ApiKey,
		}
	}
	return result
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func DatabaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = DatabaseFeedToFeed(feed)
	}
	return result
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}
