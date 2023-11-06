package video

import (
	"context"
)

// Aggregate root
type Video interface {
	Id() string
	URL() string

	SetTags(ctx context.Context, tags []tag) error

	Like(ctx context.Context, userId string) Like
	Dislike(ctx context.Context, userId string) error

	Share(userId string) Share
}
