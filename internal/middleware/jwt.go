// package middleware provides the http middlewares.
package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/emregocer/golang_project/internal/handler"
	model "github.com/emregocer/golang_project/internal/model/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type ContextKey string

const ContextUserId ContextKey = "userID"

// Jwt middleware checks the validity of the tokens provided with the Bearer Authorization header.
//
// Parameters:
//	- `h` : httprouter.Handle
//	- `jwtKey` : string
func Jwt(h httprouter.Handle, jwtKey string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		jwtKey := []byte(jwtKey)

		authPrefix := "Bearer "

		authHeader := r.Header.Get("Authorization")

		if strings.HasPrefix(authHeader, authPrefix) {
			accessToken := authHeader[len(authPrefix):]
			claims := new(model.Claims)

			if accessToken != "" {
				token, err := jwt.ParseWithClaims(accessToken, claims,
					func(token *jwt.Token) (interface{}, error) {
						return jwtKey, nil
					})

				// Add the user id to the request context.
				r = r.WithContext(context.WithValue(r.Context(), ContextUserId, claims.UserID))

				if token.Valid {
					if claims.ExpiresAt < time.Now().Unix() {
						// expired
						handler.JsonErrorResponse(w, http.StatusUnauthorized, "Not authorized")
						return
					}
				} else if validationError, ok := err.(*jwt.ValidationError); ok {
					if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
						// malformed token
						handler.JsonErrorResponse(w, http.StatusUnauthorized, "Not authorized")
						return
					} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						// validation failed
						handler.JsonErrorResponse(w, http.StatusUnauthorized, "Not authorized")
						return
					} else {
						// can't handle the token
						handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
						return
					}
				}
			} else {
				// Token is empty
				handler.JsonErrorResponse(w, http.StatusUnauthorized, "Not authorized")
				return
			}

		} else {
			// No token in the header
			handler.JsonErrorResponse(w, http.StatusUnauthorized, "Not authorized")
			return
		}

		// Handle the request with the added context
		h(w, r, ps)
	}
}
