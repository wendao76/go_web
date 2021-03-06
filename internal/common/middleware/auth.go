package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_web/internal/oauth2/model"
	"gopkg.in/oauth2.v3/generates"
	"log"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("Authorization"); token != "" {
			if passed := AuthByToken(token, ctx); !passed {
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}

//TODO 根据用户密码登录
func AuthByUsername(username string, password string, ctx *gin.Context) bool {
	return true
}

//根据token进行登录
func AuthByToken(access string, ctx *gin.Context) bool {
	token, err := jwt.ParseWithClaims(access, &model.JWTAccessClaimsLocal{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte("00000000"), nil
	})
	if err != nil {
	    log.Fatal(err.Error())
	}

	claims, ok := token.Claims.(*model.JWTAccessClaimsLocal)
	if !ok || !token.Valid {
		log.Fatal("invalid token")
		return false
	}
	fmt.Println("jwt数据:")
	fmt.Printf("Username:%s, Uid:%d", claims.Username, claims.Uid)
	ctx.Set("auth_user", token)
	return true
}

func AuthByTokenOld(access string, ctx *gin.Context) bool{
	token, err := jwt.ParseWithClaims(access, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte("00000000"), nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	cms := &model.JWTAccessClaimsLocal {
		Username : "wendao76",
		Uid: 10000,
	}
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cms)
	tokenStr, _ := testToken.SignedString([]byte("00000000"))
	log.Println("token gen")
	log.Println(tokenStr)

	claims, ok := token.Claims.(*generates.JWTAccessClaims)
	if !ok || !token.Valid {
		log.Fatal("invalid token")
		return false
	}
	fmt.Println("jwt数据:")
	fmt.Printf("token.Raw:%s, ExpiresAt:%d, Subject:%s, Audience:%srrrrrrrrrrrrrrrrrrrrrrrrrrrrrf4eedddddddddddddddd", token.Raw, claims.ExpiresAt, claims.Subject, claims.Audience)
	ctx.Set("auth_user", token)
	return true
}

//TODO 根据小程序信息进行登录
func AuthByMina(loginData interface{}) bool{
	return true
}
