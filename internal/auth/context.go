package auth

import "context"

// (1) 커스텀 컨텍스트 키 정의 (외부에 노출되지 않도록 unexported)
type contextKey string

const userContextKey contextKey = "user"

// (2) 컨텍스트에 사용자 정보를 저장하는 함수(Setter)
func WithUserContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, userContextKey, username)
}

// (3) 컨텍스트에서 사용자 정보를 가져오는 함수(Getter)
func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userContextKey).(string)
	return user, ok
}
