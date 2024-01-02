package libsession

import (
	"context"
	"time"
)

type SessionJWT struct {
	TokenID            int64 `json:"token_id"`
	UserID             int64 `json:"user_id"`
	RoleID             int64 `json:"role_id"`
	OrganizationID     int64 `json:"organization_id"`
	OrganizationRoleID int64 `json:"organization_role_id"`
}

type Session struct {
	SessionJWT
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
