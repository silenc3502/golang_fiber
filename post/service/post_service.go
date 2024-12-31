package service

import "golang_fiber/post/entity"

type PostService interface {
	Create(post *entity.Post) error
	Read(id uint) (*entity.Post, error)
	List() ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(id uint) error
}
