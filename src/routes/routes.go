package routes

import (
	"Ecom/src/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(requestRoutes *gin.Engine) {
	requestRoutes.POST("/user/signup", controllers.userSignUp())
	requestRoutes.POST("/user/login", controllers.userLogin())
	requestRoutes.POST("/admin/addProduct", controllers.adminAddProduct())
	requestRoutes.GET("/user/ProductView", controllers.userProductView())
	requestRoutes.GET("/user/search", controllers.userSearch())

}
