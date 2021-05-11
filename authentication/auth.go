package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("xkcvmxclvsdlvkn123124")

type accessClaims struct {
	jwt.StandardClaims
	UserID uint
}

type refreshClaims struct {
	jwt.StandardClaims
	UserID uint
}

//GenerateTokenPair function generates jwt token pair
func GenerateTokenPair(userID uint) ([]string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessClaims{jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 5).Unix()}, userID})
	signedAccessToken, err := accessToken.SignedString(secretKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshClaims{jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 48).Unix()}, userID})
	signedRefreshToken, err := refreshToken.SignedString(secretKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return []string{signedAccessToken, signedRefreshToken}, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil

}

func RefreshToken(token string) (string, error) {
	decoded, err := ValidateToken(token)
	if err != nil {
		fmt.Println("error line 66")
		return "", fmt.Errorf("authentication error")
	}
	pair, err := GenerateTokenPair(uint(decoded["UserID"].(float64)))
	if err != nil {
		fmt.Println("error line 72")
		return "", fmt.Errorf("authentication error")
	}
	return pair[0], nil
}
