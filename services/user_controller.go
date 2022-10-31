package services

import (
	"ThirdEssentials/db"
	"ThirdEssentials/services/core/users"
	"ThirdEssentials/services/payload"
	"ThirdEssentials/services/response"
	"ThirdEssentials/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	request := new(payload.RegisterUser)

	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{
			Message: "Bad Request",
		})
	}

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	tx := db.Begin()
	userRes, err := users.UserRegistration(tx, request)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func LoginUser(c *gin.Context) {

	request := new(payload.LoginUser)

	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{
			Message: "Bad Request",
		})
	}

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	tx := db.Begin()
	userRes, err := users.LoginUser(tx, request)

	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func GetAllUsers(c *gin.Context) {

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	userRes, err := users.GetUsersRecords(db)
	if err != nil {
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	db.Close()
	c.JSON(http.StatusOK, userRes)

}

func Logoutuser(c *gin.Context) {

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Message: "Database connection Error"})
	}

	isAuthenticated, userID, token := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	tx := db.Begin()
	err = users.UpdateToken(tx, token, userID)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, response.Message{Message: "Could not logout user"})
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, response.Message{Message: "Logged Out successfully!"})
}

func FetchUserLoginActivities(c *gin.Context) {

	id := c.Params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{Message: "Invalid user ID"})
	}

	isAuthenticated, superAdminID, _ := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !users.CheckSuperAdmin(db, superAdminID) {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "Not authorized to perform this action",
		})
		return

	}

	tx := db.Begin()
	res, err := users.FetchUserLoginActivity(tx, userID)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, res)
}
