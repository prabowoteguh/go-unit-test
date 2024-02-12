package service

import (
	"fmt"
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)
	product, err := productService.GetOneProduct("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

// func TestProductServiceGetOneProduct(t *testing.T) {
// 	product := entity.Product{Id: "2", Name: "Kaca mata"}

// 	productRepository.Mock.On("FindById", "2").Return(product)

// 	result, err := productService.GetOneProduct("2")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
// 	assert.Equal(t, product.Name, result.Name, "result has to be 'Kaca mata'")
// 	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
// }

func TestProductServiceGetAllProduct(t *testing.T) {
	product := []entity.Product{{Id: "2", Name: "Kaca mata"}, {Id: "3", Name: "Jam tangan"}}

	productRepository.Mock.On("FindAll").Return(product)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(product), len(result), "result has to be '2'")

	fmt.Println(result)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	for i, v := range result {
		assert.Equal(t, product[i].Id, v.Id, "Id should be equal"+product[i].Id)
		assert.Equal(t, product[i].Name, v.Name, "Name should be equal"+product[i].Name)
		assert.Equal(t, product[i], v, "result has to be a product data with id"+product[i].Id)
	}
}

func TestFailProductServiceGetAllProducts(t *testing.T) {
	products := []entity.Product{
		{Id: "1", Name: "Glasses"},
		{Id: "2", Name: "Sunglasses"},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, _ := productService.GetAllProduct()

	assert.NotEqual(t, len(products), len(result)+1, "failed to get all products")

	for i, v := range result {
		assert.Equal(t, products[i].Name, v.Name, "Name should be equal"+products[i].Name)
	}
}
