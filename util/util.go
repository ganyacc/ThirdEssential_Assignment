package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	database "ThirdEssentials/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const Secret = "XCFWQFSFSFUR"

func HashPass(pwd string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckHash(hash string, pwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func Authorize(c *gin.Context) (bool, int, string) {

	header := c.Request.Header

	auth := header.Get("Authorization")

	if len(strings.Trim(auth, " ")) == 0 {
		fmt.Println("--------------Invalid Auth 1")
		return false, 0, ""
	}

	headerComponent := strings.Split(auth, " ")
	if len(headerComponent) != 2 {
		fmt.Println("--------------Invalid Auth 2")
		return false, 0, ""
	}

	tokenType := headerComponent[0]
	tokenString := headerComponent[1]

	if tokenType != "JWT" {
		fmt.Println("--------------Invalid Auth 3")
		return false, 0, ""
	}

	// database.Token()
	db, err := database.InitDb()
	if err != nil {
		return false, 0, ""
	}

	var tokenObj database.Token

	db.Model(database.Token{}).Where(database.Token{
		Token: tokenString,
	}).Find(&tokenObj)

	diff := tokenObj.Expiry.Sub(time.Now())
	fmt.Println("=====>>>>", tokenObj.Expiry, time.Now(), diff.Seconds())

	if diff.Seconds() < 0 {
		return false, 0, ""
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(Secret), nil
	})

	if err != nil {
		return false, 0, ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok == false {
		return false, 0, ""
	}

	userID := claims["user_id"]

	userId := int(userID.(float64))

	if userId == 0 {
		return false, 0, ""
	}

	return true, userId, tokenString
}
