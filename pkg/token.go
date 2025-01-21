package pkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func TokenValid(c *gin.Context) error {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}

		return []byte(viper.GetString("SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("access_token")
	if err != nil {
		return "", err
	}
	if token != "" {
		return token, nil
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}

	return "", nil
}

func ExtractTokenId(c *gin.Context) (any, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(viper.GetString("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, err
	}

	fmt.Printf("Token claims: %+v\n", claims)

	return claims["subject"], nil
}
