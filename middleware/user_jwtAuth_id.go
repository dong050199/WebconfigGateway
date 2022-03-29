package middleware

// import (
// 	handler "SQLite_Repo_Pattern/controller/user_handler"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"strconv"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gorilla/mux"
// )

// func AuthenticationWithId(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//do somethings
// 		fmt.Printf("Check token by middleware")
// 		tokenHeader := r.Header.Get("Authorization")

// 		if tokenHeader == "" {
// 			handler.ResponseErr(w, http.StatusForbidden)
// 			return
// 		}

// 		splitted := strings.Split(tokenHeader, " ") // Bearer jwt_token
// 		if len(splitted) != 2 {
// 			handler.ResponseErr(w, http.StatusForbidden)
// 			return
// 		}

// 		tokenPart := splitted[1]
// 		tk := &Claims{}

// 		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil {
// 			fmt.Println(err)
// 			handler.ResponseErr(w, http.StatusInternalServerError)
// 			return
// 		}

// 		if token.Valid {
// 			handler.ResponseOk(w, token.Claims)
// 		}

// 		w.Header().Set("Content-Type", "applycation/json")
// 		params := mux.Vars(r)
// 		id, _ := strconv.Atoi(params["id"])
// 		ctx := context.WithValue(r.Context(), "id", id)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
