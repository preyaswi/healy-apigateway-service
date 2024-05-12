package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}

	return header
}
func ExtractUserIDFromToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}
	fmt.Println(claims.Id,"id is ")

	return claims.Id, claims.Email, nil

}
