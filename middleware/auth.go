package middleware

import (
	"context"
	"errors"
	"net/http"

	"app/models"
)

func RequireLogin(next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionID := cookie.Value
		user := &models.AuthenticatedUser{SessionID: sessionID}

		userErr := user.GetUserBySessionID()
		if userErr != nil {
			http.Error(w, userErr.Error(), userErr.Code)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next(w, r)
	}
}

func GetAuthenticatedUser(r *http.Request) (*models.AuthenticatedUser, error) {
	user := r.Context().Value("user")
	authenticatedUser, ok := user.(*models.AuthenticatedUser)
	if !ok {
		return nil, errors.New("token not found")
	}
	return authenticatedUser, nil
}
