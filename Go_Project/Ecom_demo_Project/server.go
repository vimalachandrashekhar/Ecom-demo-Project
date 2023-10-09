package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/config"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/controllers"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/routes"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/services"
	"go.mongodb.org/mongo-driver/mongo"

)

var (
	mongoClient *mongo.Client
	err         error
	server      *gin.Engine
)

func main() {

	server = gin.Default()
	InitializeDatabase()
	InitializeProducts()
	InitializeUsers()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx)

	server.Run(":4000")

}

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to Database %v", err)
	} else {
		fmt.Println("Connected to Database")
	}
}

func InitializeProducts() {
	productCollection := config.GetCollection(mongoClient, "ecom", "products")
	productService := services.InitProductService(productCollection)
	productController := controllers.InitProductController(productService)
	routes.ProductRoutes(server, *productController)
}

func InitializeUsers() {
	userCollection := config.GetCollection(mongoClient, "ecom", "users")
	userService := services.InitUserService(userCollection)
	userController := controllers.InitUserController(userService)
	routes.UserRoutes(server, *userController)
}
