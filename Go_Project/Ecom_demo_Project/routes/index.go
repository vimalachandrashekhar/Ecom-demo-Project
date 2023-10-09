package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/controllers"
)

func ProductRoutes(r *gin.Engine, pc controllers.ProductController) {
	pRoute := r.Group("/ecom/product")
	pRoute.POST("/addproduct", pc.AddProduct)
	pRoute.GET("/getproduct/:id", pc.GetProductById)
	pRoute.GET("/getproducts/:name", pc.GetProductsByName)
}

func UserRoutes(r *gin.Engine, uc controllers.UserController) {
	uRoute := r.Group("/ecom/user")
	uRoute.POST("/register", uc.Register)
}
