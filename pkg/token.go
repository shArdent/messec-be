package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(email, id string) (string, error) {
	claims := jwt.MapClaims{
		"issuer":  "mesecret",
		"email":   email,
		"subject": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(viper.GetString("SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
