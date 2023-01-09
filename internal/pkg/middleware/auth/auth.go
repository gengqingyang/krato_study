/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-16 15:30:15
 * @LastEditTime: 2022-12-16 16:54:35
 * @Description:
 */
package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt"
)

func JwtAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				tokenString := tr.RequestHeader().Get("Authorization")
				fmt.Println("Authorization is :", tokenString)

				auths := strings.SplitN(tokenString, " ", 2)

				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, fmt.Errorf("miss token string")
				}
				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(secret), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["username"])
				} else {
					return nil, err
				}
			}
			return handler(ctx, req)
		}
	}
}

func GenerateToken(username, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("GenerateToken err is:", err)
		return "", err
	}

	return tokenString, nil
}
