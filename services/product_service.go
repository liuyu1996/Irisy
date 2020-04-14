package services

import (
	"product/datamodels"
	"product/repositories"
)

type IProductService interface {
	GetProductByID(int64)(*datamodels.Product, error)
	GetAllProduct()([]*datamodels.Product, error)
	DeleteProductByID(int64)bool
	InsertProduct(product *datamodels.Product)(int64, error)
	UpdateProduct(product *datamodels.Product)error
}

type ProductService struct {
	productRepostory repositories.IProduct
}

func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{repository}
}

func (p *ProductService) GetProductByID(productID int64)(*datamodels.Product, error) {
	return p.productRepostory.SelectByKey(productID)
}

func (p *ProductService) GetAllProduct()([]*datamodels.Product, error) {
	return p.productRepostory.SelectAll()
}

func (p *ProductService) DeleteProductByID(productID int64)bool {
	return p.productRepostory.Delete(productID)
}

func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepostory.Insert(product)
}

func (p *ProductService) UpdateProduct(product *datamodels.Product)error {
	return p.productRepostory.Update(product)
}