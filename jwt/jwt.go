package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCredential struct {
	TokenSecret         string
	ExpiredToken        int
	RefreshTokenSecret  string
	ExpiredRefreshToken int
}

type CustomClaims struct {
	jwt.StandardClaims
	Payload string      `json:"payload"`
}

func (cred JwtCredential) GetToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(cred.ExpiredToken) * time.Hour).Unix()
	//unixTimeUtc := time.Unix(expirationTime, 0)
	//unitTimeInRFC3339 := unixTimeUtc.UTC().Format(time.RFC3339)

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.TokenSecret))

	return token, expirationTime, err
}

func (cred JwtCredential) GetRefreshToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(cred.ExpiredRefreshToken) * time.Hour).Unix()
	//unixTimeUtc := time.Unix(expirationTime, 0)
	//unitTimeInRFC3339 := unixTimeUtc.UTC().Format(time.RFC3339)

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.RefreshTokenSecret))

	return token, expirationTime, err
}
