package routes

import (
	"base-project-api/controllers"
	middlewares "base-project-api/server/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	main.POST("/login", controllers.Login)
	main.POST("/user", controllers.CreateUser)
	{
		routers := main.Group("/", middlewares.AuthJwt())
		{
			routers.GET("/", controllers.HellowControllers)
			routers.POST("/seller", controllers.CreateSeller)
			routers.GET("/seller/:id", controllers.GetSellerById)
			routers.GET("/seller", controllers.GetAllSellers)
			routers.DELETE("/seller/:id", controllers.DeleteSeller)
			routers.GET("/user", controllers.GetUserInfo)
			routers.GET("/company", controllers.GetCompany)
			routers.POST("/company", controllers.CreateCompany)
			routers.DELETE("/company/:id", controllers.DeleteCompany)
			routers.PUT("/company/", controllers.UpdateCompany)
			routers.GET("/product", controllers.GetAllProducts)
			routers.GET("/product/:id", controllers.GetProductByID)
			routers.POST("/product", controllers.CreateProduct)
			routers.PUT("/product", controllers.UpdateProduct)
			routers.DELETE("/product/:id", controllers.DeleteProduct)
			routers.GET("/address/:id", controllers.GetAddresses)
			routers.POST("/address", controllers.CreateAddress)
			routers.PUT("/address", controllers.UpdateAddress)
			routers.DELETE("/address/:id", controllers.DeleteAddress)
			routers.GET("/order", controllers.GetOrders)
			routers.POST("/order", controllers.CreateOrder)
			routers.PUT("/order", controllers.UpdateOrder)
			routers.DELETE("/order/:id", controllers.DeleteOrder)
			routers.POST("/orderitems", controllers.InsertItemsOrder)
			routers.PUT("/orderitems", controllers.UpdateOrderItems)
			routers.DELETE("/orderitems/:id", controllers.DeleteOrderItems)
			routers.GET("/orderitems/:id", controllers.GetAllItemsInOrder)

		}
	}
	return router
}
