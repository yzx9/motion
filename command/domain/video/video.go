package video

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/yzx9/motion/domain/video/adapter/qiniu"
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

func Upload(ctx context.Context,
	name, cratorId string, f *multipart.FileHeader,
) (Video, error) {
	if err := qiniu.Upload(f); err != nil {
		return nil, err // TODO: wrap err
	}
	return nil, fmt.Errorf("not implement") // TODO
}
