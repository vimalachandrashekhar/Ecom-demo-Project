package services

import (
	"context"
	"fmt"
	"time"

	"github.com/vimalachandrashekhar/Ecom-demo-Project/entities"
	"github.com/vimalachandrashekhar/Ecom-demo-Project/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {
	return &ProductService{Product: collection}
}

func (prod *ProductService) AddProduct(p *entities.Product) (string, error) {
	p.Id = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt
	_, err := prod.Product.InsertOne(context.Background(), p)
	if err != nil {
		return "", err
	} else {
		return "Record Inserted Successfully", nil
	}
}

func (prod *ProductService) GetProductById(id primitive.ObjectID) (*entities.Product, error) {

	ctx := context.Background()
	var product entities.Product
	err := prod.Product.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (prod *ProductService) SearchProducts(name string) ([]*entities.Product, error) {
	var products []*entities.Product
	cursor, err := prod.Product.Find(context.TODO(), bson.M{"name": name})
	if err != nil {
		return nil, err
	} else {
		fmt.Println(cursor)
		for cursor.Next(context.TODO()) {
			product := &entities.Product{}
			err := cursor.Decode(product)

			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
		if err := cursor.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}
}
