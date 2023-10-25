package video

import "context"

// aggrete
type Share interface {
	Token() string

	Touch(ctx context.Context, ip string)
	TouchByUser(ctx context.Context, ip, userId string)
}
