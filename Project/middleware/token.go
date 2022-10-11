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
		//user := UserInfo{
		//	Account:   "账号",
		//	Password:  "密码",
		//	AccountId: "账号id",
		//}
		//tokenString, _err := GenerateToken(user)
		//fmt.Println("tokenString", tokenString)
		_claims, _err := ParseToken(_tokenStr)
		if _err != nil {
			response(ctx, newResponse(40101, _err.Error(), ""))
			return
		} else {
			fmt.Println("=========", _claims)
		}
		ctx.Set("account", _claims.Account)
		ctx.Set("account_id", _claims.AccountId)
		ctx.Next()
	}
}

type UserInfo struct {
	Account   string `form:"account"`
	Password  string `form:"password"`
	AccountId string `form:"account_id"`
}

type MyClaims struct {
	UserInfo
	jwt.RegisteredClaims
}

var MySecret = []byte("手写的从前")

func GenerateToken(user UserInfo) (tokenString string, err error) {
	claim := MyClaims{
		UserInfo: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(240 * time.Hour * time.Duration(1))),
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
	fmt.Println(tokens)
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("无效令牌")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("令牌已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("令牌尚未激活")
			} else {
				return nil, errors.New("无效令牌")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效令牌")
}
