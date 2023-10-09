package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/entities"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) AddProduct(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	if res, err := p.ProductService.AddProduct(&product); err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, res)
	}
}

func (p ProductController) GetProductById(c *gin.Context) {

	productId := c.Param("id")
	pid, err := primitive.ObjectIDFromHex(productId)
	product, err := p.ProductService.GetProductById(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, product)
}

func (p ProductController) GetProductsByName(c *gin.Context) {
	name := c.Param("name")
	products, err := p.ProductService.SearchProducts(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, products)
}
