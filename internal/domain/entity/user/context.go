package user

import (
	"context"
)

type Key int

var (
	UserKey Key = 0
)

func ExtractFromCtx(ctx context.Context) *User {
	u, ok := ctx.Value(UserKey).(*User)
	if !ok {
		return nil
	}
	return u
}

func PutToCtx(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, UserKey, u)
}
