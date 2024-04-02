package http

import (
	"chatapp/internal/shop"
	"chatapp/internal/user"

	"github.com/gin-gonic/gin"
)
var r *gin.Engine
func InitRouter(userHanler *user.Handler,shopHandler *shop.Handler) {
	r = gin.Default()
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/signup",userHanler.Signup)
		userRoutes.POST("/login",userHanler.Login)
	}
	shopRoutes := r.Group("/shop")
	{
		shopRoutes.POST("/addProduct",shopHandler.AddProduct)
		shopRoutes.GET("/getAllProducts",shopHandler.GetAllProducts)

	}
}
func Start(addr string)error{
	return r.Run(addr)
}