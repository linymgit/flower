package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

const secretKey = "forrily"

func GenJwt(claims jwt.MapClaims) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err = token.SignedString([]byte(secretKey))
	if err != nil {
		// TODO
	}
	return
}

func ParseJwt(passToken string) (claims jwt.MapClaims, ok bool, err error) {
	token, err := jwt.Parse(passToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		// TODO
		return
	}
	if token == nil {
		// TODO
		return
	}
	if claims, ok = token.Claims.(jwt.MapClaims); !ok || claims.Valid() != nil {
		// TODO
		return
	}
	return
}
