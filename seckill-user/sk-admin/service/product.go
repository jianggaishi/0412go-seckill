package service

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"log"
	"seckill/sk-admin/model"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProductList() ([]gorose.Data, error)
}

type ProductServiceMiddleware func(ProductService) ProductService

type ProductServiceImpl struct {
}

func (p ProductServiceImpl) CreateProduct(product *model.Product) error {
	productEntity := model.NewProductModel()
	fmt.Println("我要开始创建商品了")
	err := productEntity.CreateProduct(product)
	if err != nil {
		log.Printf("ProductEntity.CreateProduct, err : %v", err)
		fmt.Println("创建商品的错误是：", err)
		return err
	}
	return nil
}

func (p ProductServiceImpl) GetProductList() ([]gorose.Data, error) {
	productEntity := model.NewProductModel()
	productList, err := productEntity.GetProductList()
	if err != nil {
		log.Printf("ProductEntity.CreateProduct, err : %v", err)
		return nil, err
	}
	return productList, nil
}
