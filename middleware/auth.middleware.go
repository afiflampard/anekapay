package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ResponseWithError(w, http.StatusBadRequest, "Authorization required")
			return
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			ResponseWithError(w, http.StatusBadRequest, "Authorization required")
			return
		}
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Something Wrong")
			}
			return []byte("I love Malang"), nil
		})
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			fmt.Println(err)
			ResponseWithError(w, http.StatusBadRequest, "Unauthorization")
		}

		ctx := context.WithValue(r.Context(), "userToken", claims["id"])
		userId := strconv.FormatFloat(claims["id"].(float64), 'g', 1, 64)
		r.Header.Set("user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ResponseWithError(w http.ResponseWriter, code int, message interface{}) {
	payload := map[string]interface{}{"error": message}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
