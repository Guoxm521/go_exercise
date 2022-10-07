package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_tokenStr := ctx.GetHeader("Authorization")
		if _tokenStr == "" {
			response(ctx, newResponse(40101, "令牌不合法", ""))
			return
		}
		user := userInfo{
			Account:   "账号",
			Password:  "密码",
			AccountId: "账号id",
		}
		tokenString, err := GenerateToken(user)
		fmt.Println("===============", tokenString)
		fmt.Println("================err", err)
		_d, _ := ParseToken(tokenString)
		fmt.Println("=========", _d.AccountId)
		ctx.Next()
	}
}

type userInfo struct {
	Account   string `form:"account"`
	Password  string `form:"password"`
	AccountId string `form:"account_id"`
}

type MyClaims struct {
	userInfo
	jwt.RegisteredClaims
}

var MySecret = []byte("手写的从前")

func GenerateToken(user userInfo) (tokenString string, err error) {
	claim := MyClaims{
		userInfo: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("手写的从前"), nil
	}
}
func ParseToken(tokens string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
