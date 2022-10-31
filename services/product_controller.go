package services

import (
	"ThirdEssentials/db"
	"ThirdEssentials/services/core/product"
	"ThirdEssentials/services/core/users"
	"ThirdEssentials/services/payload"
	"ThirdEssentials/services/response"
	"ThirdEssentials/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {

	request := new(payload.Product)

	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{
			Message: "Bad Request",
		})
		return
	}

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	isAuthenticated, userID, _ := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	tx := db.Begin()
	userRes, err := product.AddProduct(tx, userID, request)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func GetAllProductbyId(c *gin.Context) {

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	isAuthenticated, userID, _ := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	tx := db.Begin()
	userRes, err := product.GetAllProducts(tx, userID)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func GetAllProducts(c *gin.Context) {

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	isAuthenticated, userID, _ := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	tx := db.Begin()
	userRes, err := product.GetAllProducts(tx, userID)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func UpdateProduct(c *gin.Context) {

	id := c.Params.ByName("id")
	pId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{Message: "counld not parse id"})
	}

	request := new(payload.Product)
	err = c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{
			Message: "Bad Request",
		})
		return
	}

	isAuthenticated, userID, _ := util.Authorize(c)
	fmt.Println("USER IS ==!!!!==?>>>>>>", isAuthenticated, userID)

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

	tx := db.Begin()
	userRes, err := product.UpdateProduct(tx, userID, pId, request)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, userRes)
}

func DeleteProduct(c *gin.Context) {

	id := c.Params.ByName("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{
			Message: "counld not parse id",
		})
	}

	db, err := db.InitDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	isAuthenticated, userId, _ := util.Authorize(c)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "UnAuthorized",
		})

		return
	}

	tx := db.Begin()
	err = product.DeleteProductbyID(tx, productID, userId)
	if err != nil {
		tx.Rollback()
		db.Close()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()
	db.Close()
	c.JSON(http.StatusOK, response.Message{
		Message: "Success",
	})
}

func GetUserProductActivity(c *gin.Context) {

	id := c.Params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Message{"Invalid user ID"})
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

	fmt.Println("=============================================")
	if !users.CheckSuperAdmin(db, superAdminID) {
		c.JSON(http.StatusUnauthorized, response.Message{
			Message: "Not authorized to perform this action",
		})
		return
	}

	tx := db.Begin()
	res, err := product.GetUserProductActivity(tx, userID)
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
