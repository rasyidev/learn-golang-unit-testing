package service

import (
	"errors"
	"learn-go-unit-test/entity"
	"learn-go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)

	if category == nil {
		return nil, errors.New("kategori tidak ditemukan")
	} else {
		return category, nil
	}
}
