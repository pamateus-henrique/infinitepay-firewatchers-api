package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJWT(username string, id int) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"user_id": id, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), })

	secret, ok := os.LookupEnv("JWT_SECRET");
	
	if !ok {
		log.Fatal("Problems getting jwt_secret from env")
	}

	jwt, err := token.SignedString([]byte(secret))
	
	if err != nil {
		return "", err
	}

	return jwt,nil

	}