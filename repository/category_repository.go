package repository

import "learn-go-unit-test/entity"

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}
