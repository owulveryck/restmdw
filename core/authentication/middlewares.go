package authentication

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func GetToken(req *http.Request) (string, error) {
	var err error
	if token := req.Header.Get("Authorization"); token != "" {
		// Should be a bearer token
		if len(token) > 6 && strings.ToUpper(token[0:6]) == "BEARER" {
			token = token[7:]
			return token, nil
		} else {
			err = fmt.Errorf("Cannot extract token from request: not a Bearer authentication)")
		}
	} else {
		err = fmt.Errorf("Cannot extract token from request: No authentication header found")

	}
	return "", err
}

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
