package util

/**
*@Description:
*@Author: BZ
*@date: 2023/11/4 11:56
*@Version: V1.0
 */

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	UserID     int
	Username   string
	GrantScope string
	jwt.RegisteredClaims
}

// 签名密钥
const (
	SIGN_kEY             = "blue-princess-motion-video"
	ACCESS_USER          = "user"
	ACCESS_VISITOR       = "visitor"
	ACCESS_ADMINISTRATOR = "administrator"
)

func GenerateToken(userId int, userName string) (string, error) {
	claim := MyCustomClaims{
		UserID:     userId,
		Username:   userName,
		GrantScope: ACCESS_USER,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "motion_Server",                                    // 签发者
			Subject:   "user",                                             // 签发对象agent
			Audience:  jwt.ClaimStrings{"ANDROID_APP", "IOS_APP", "WEB"},  //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),    //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			ID:        RandomString(10),                                   // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(SIGN_kEY))
	return token, err
}

func ParseToken(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGN_kEY), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
