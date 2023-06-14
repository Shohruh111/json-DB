package storage

import (
	"app/models"
)

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.User, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(*models.UpdateUser) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}

type ProductRepoI interface{
	Create(*models.CreateProduct) (*models.Product, error)
	GetById(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(*models.UpdateProduct) (*models.Product, error)
	Delete(*models.ProductPrimaryKey) error
}
