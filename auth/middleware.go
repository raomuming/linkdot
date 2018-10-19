package auth

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/raomuming/linkdot/model"
	"github.com/raomuming/linkdot/utils"
)

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"Id":   user.Id,
	})

	return token.SignedString([]byte("secret"))
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			utils.ResponseWithJson(w, http.StatusUnauthorized,
				utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					utils.ResponseWithJson(w, http.StatusUnauthorized,
						utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
				}
				return []byte("secret"), nil
			})

			if !token.Valid {
				utils.ResponseWithJson(w, http.StatusUnauthorized,
					utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
			} else {
				// https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					ctx := context.WithValue(r.Context(), "UserName", claims["username"])
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.ResponseWithJson(w, http.StatusUnauthorized,
						utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
				}
			}
		}
	})
}
