package util

import (
	"time"

	"github.com/hechen0210/utils/helper"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type Jwt struct {
	Exception []string
	Secret    string
}

func NewJwt(info Jwt) *jwt.Middleware {
	return jwt.New(jwt.Config{
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil || helper.Contains(info.Exception, ctx.Path()) {
				return
			}
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(map[string]interface{}{
				"code":    "401",
				"message": err.Error(),
			})
		},
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(info.Secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func CreateToken(info map[string]interface{}, secret string) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range info {
		claims[k] = v
	}
	claims["time"] = time.Now().Unix()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("My Secret"))
}
