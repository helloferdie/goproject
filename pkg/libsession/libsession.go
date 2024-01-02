package libsession

import (
	"context"
	"time"
)

type Session struct {
	UserID    int64
	AccountID int64
	TokenID   int64
	IPAddress string
	Language  string
	Timezone  *time.Location
}

func NewContext(ctx context.Context, s *Session) context.Context {
	return context.WithValue(ctx, "session", s)
}

func FromContext(ctx context.Context) (*Session, bool) {
	s, ok := ctx.Value("session").(*Session)
	return s, ok
}
