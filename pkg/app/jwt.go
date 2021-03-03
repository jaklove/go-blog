package app

import (
	"github.com/dgrijalva/jwt-go"
	"go-blog/global"
	"go-blog/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return global.JWTSetting.Secret
}

type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpireAt  int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func GenerateToken(appkey, appScret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	Claims := Claims{
		AppKey:    util.EncodeMD5(appkey),
		AppSecret: util.EncodeMD5(appScret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, Claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string)(*Claims,error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}