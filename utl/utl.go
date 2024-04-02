package utl

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(pass []byte) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func ComparePasswords(loginPass string, hashPass string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(loginPass)); err != nil {
		return err
	}
	return nil
}
func GenerateJWT(username string, id string) (string, error) {
	//generate JWT token and returning it
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = id
	claims["exp"] = time.Now().Add(60 * time.Second * 24 * 30).Unix()
	claims["authorized"] = true
	var sampleSecretKey = []byte("secretkey")
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
