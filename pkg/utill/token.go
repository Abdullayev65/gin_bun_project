package utill

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type TokenJWT struct {
	signingKey     []byte
	ExpiryInterval time.Duration
}

func NewToken(signingKey string, expiryInterval time.Duration) *TokenJWT {
	return &TokenJWT{signingKey: []byte(signingKey), ExpiryInterval: expiryInterval}
}

func (t *TokenJWT) Generate(sub int64) (string, error) {
	return t.generate(strconv.FormatInt(sub, 10))
}

func (t *TokenJWT) Parse(tokenStr string) (int64, error) {
	subStr, err := t.parse(tokenStr)
	if err != nil {
		return 0, err
	}
	sub, err := strconv.ParseInt(subStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return sub, nil
}

func (t *TokenJWT) generate(sub string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(t.ExpiryInterval).Unix(),
		Subject:   sub,
	})
	return token.SignedString(t.signingKey)
}

func (t *TokenJWT) parse(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return t.signingKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("token claims are not type of  jwt.StandardClaims")
	}
	return claims.Subject, nil
}