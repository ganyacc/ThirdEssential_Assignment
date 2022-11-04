package services

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.POST("/user/register", RegisterUser)
	r.POST("/user/login", LoginUser)
	r.POST("/user/logout", Logoutuser)

	r.POST("/product", AddProduct)
	r.PUT("/product/:id", UpdateProduct)
	r.GET("/products", GetAllProducts)
	r.DELETE("/product/:id", DeleteProduct)

	//super admin apis
	r.GET("/users", GetAllUsers)

	r.GET("/user/:id/product/activities", GetUserProductActivity)
	r.GET("/user/:id/activities", FetchUserLoginActivities)

}
