package users

import (
	database "ThirdEssentials/db"
	"ThirdEssentials/services/payload"
	"ThirdEssentials/services/response"
	"ThirdEssentials/util"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

func UserRegistration(db *gorm.DB, request *payload.RegisterUser) (*database.Users, error) {

	password, err := util.HashPass(request.Password)

	if err != nil {
		return nil, err
	}

	newUser := database.Users{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Role:     request.Role,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	err = db.Create(&newUser).Error

	if err != nil {
		return nil, err
	}

	return &newUser, nil

}

func LoginUser(db *gorm.DB, request *payload.LoginUser) (*response.Token, error) {

	var user database.Users

	notfound := db.Model(database.Users{}).Where(database.Users{
		Email: request.Email,
	}).Find(&user).RecordNotFound()

	if notfound {
		return nil, errors.New("user not found")
	}

	isCorrectPass := util.CheckHash(user.Password, request.Password)
	if !isCorrectPass {
		return nil, errors.New("incorrect Password")
	}

	db.Create(&database.UserLoginActivity{
		UserID:     user.Id,
		LogintTime: time.Now(),
	})

	token, err := createTokenForLoginForUser(db, user.Id)
	if err != nil {
		return nil, err
	}

	return &response.Token{
		Token: token.Token,
	}, nil
}

func createTokenForLoginForUser(db *gorm.DB, userId int) (*database.Token, error) {

	jwtClaims := jwt.MapClaims{}

	jwtClaims["user_id"] = userId
	jwtClaims["time"] = time.Now().String()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	jwtTokenString, err := at.SignedString([]byte(util.Secret))
	if err != nil {
		return nil, err
	}

	expiryTime := time.Now().Add(time.Hour * 24)

	token := database.Token{
		Token:  jwtTokenString,
		Expiry: expiryTime,
	}

	tokenCreateResult := db.Create(&token)
	if tokenCreateResult.Error != nil {

		return nil, tokenCreateResult.Error
	}

	return &token, nil
}

func GetUsersRecords(db *gorm.DB) ([]database.Users, error) {
	var users []database.Users
	if err := db.Model(&database.Users{}).Where(&database.Users{
		Role: "User",
	}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func FetchUserLoginActivity(db *gorm.DB, userID int) ([]database.UserLoginActivity, error) {
	var activies []database.UserLoginActivity

	if err := db.Model(&activies).Where(&database.UserLoginActivity{
		UserID: userID,
	}).Find(&activies).Error; err != nil {
		return nil, err
	}

	return activies, nil
}

func CheckSuperAdmin(db *gorm.DB, superAdminID int) bool {
	var user database.Users

	err := db.Debug().Model(database.Users{}).Where(database.Users{
		Id: superAdminID,
	}).Find(&user).Error

	if err != nil {
		return false
	}

	fmt.Println("------------------>>>>", user)

	if user.Role != "SuperAdmin" {
		return false
	}

	return true
}

func UpdateToken(db *gorm.DB, token string, userId int) error {

	err := db.Debug().Model(&database.Token{}).Where(&database.Token{
		Token: token,
	}).Update(&database.Token{
		Token:  token,
		Expiry: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&database.UserLoginActivity{
		UserID:     userId,
		LogOutTime: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return nil
}
