package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var userKey = &struct {
	key string
}{"user_id"}

func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			send(w, http.StatusUnauthorized, Reason{"fail to parse token"})
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			send(w, http.StatusUnauthorized, Reason{"fail to validate token"})
			return
		}

		user, ok := claims["user_id"].(string)
		if !ok {
			send(w, http.StatusUnauthorized, Reason{"fail to parse user_id from token"})
			return
		}
		if user == "" {
			send(w, http.StatusUnauthorized, Reason{"invalid user_id"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, user)

		log.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("user_id", user)
		})

		r = r.WithContext(ctx)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

//GetUser get user from context
func GetUser(ctx context.Context) string {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return ""
	}
	user, ok := claims["user_id"].(string)
	if user == "" || !ok {
		return ""
	}
	return user
}

//GenerateTokenCookie generate cookie which includes JWT
func GenerateTokenCookie(userID string) (*http.Cookie, error) {
	token, err := GenerateToken(userID)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
	}
	return cookie, nil
}

//GenerateToken generate token for user auth
func GenerateToken(userID string) (string, error) {
	claims := make(map[string]interface{})

	claims["user_id"] = userID
	jwtauth.SetExpiryIn(claims, time.Duration(time.Hour*1))
	jwtauth.SetIssuedNow(claims)

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
