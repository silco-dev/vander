package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/silco-dev/vander/structs"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *API) Authenticate(h http.Handler) http.Handler { // Middleware for authenticating
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, i := range a.NonSensitiveEndpoints {
			if strings.HasPrefix(r.URL.Path, i) {
				h.ServeHTTP(w, r)
				return
			}
		}

		var APIUser *structs.User
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" { // No token at all
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(Exception{Error: "Missing auth token"})
			return
		}

		if strings.HasPrefix(tokenString, "Bearer") {
			splitToken := strings.Split(tokenString, "Bearer ")
			tokenString = splitToken[1]
		}

		user, err := db.GetUser(tokenString)

		if err != nil {
			if err != mongo.ErrNoDocuments {
				log.Println("An error occured!", err)
			}
		}

		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(Exception{Error: "Auth token invalid"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(a.Secret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if APIUser.Info.Name == claims["name"] && APIUser.Info.Contact == claims["contact"] {
				if APIUser.Info.Name == "" { // No Match
					w.WriteHeader(http.StatusUnauthorized)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					json.NewEncoder(w).Encode(Exception{Error: "Auth token invalid"})
					return
				}

				if !APIUser.Enabled { // The account was disabled.
					w.WriteHeader(http.StatusUnauthorized)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					json.NewEncoder(w).Encode(Exception{Error: "Your account has been disabled"})
					return
				}

				h.ServeHTTP(w, r)

			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				json.NewEncoder(w).Encode(Exception{Error: "Auth token data mismatch"})
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Auth token invalid"})
		}

	})
}
