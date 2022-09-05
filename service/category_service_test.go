package service

import (
	"learn-go-unit-test/entity"
	"learn-go-unit-test/helper"
	"learn-go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindByID", "2").Return(category)
	res, err := categoryService.Get("2")
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)
}

func BenchmarkHelloRasyidev(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloRasyidev("Rasyidev")
	}
}

func BenchmarkHelloTaeri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloRasyidev("Kim Taeri Beauty")
	}
}

/*BENCHMARK
$ go test -v -run AplikasiNone -bench=.
goos: windows
goarch: amd64
pkg: learn-go-unit-test/service
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkHelloRasyidev
BenchmarkHelloRasyidev-8        91282789                12.68 ns/op
BenchmarkHelloTaeri
BenchmarkHelloTaeri-8           93429685                12.76 ns/op
PASS
ok      learn-go-unit-test/service      4.293s
*/
