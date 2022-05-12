package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GodKimba/cuddly-golang-server/internal/users"
	"github.com/GodKimba/cuddly-golang-server/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt token
			toeknStr := header
			username, err := jwt.ParseToken(toeknStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Create user and check if user exists in db
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)
			// put in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// Calling the next with the new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Finds the user from the context. REQUIRES Middleware to run
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
