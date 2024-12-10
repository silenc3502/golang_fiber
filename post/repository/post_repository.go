package repository

import (
	"golang_fiber/post/entity"
	"gorm.io/gorm"
)

// PostRepository는 게시글에 대한 CRUD 작업을 정의하는 인터페이스입니다.
type PostRepository interface {
	Create(post *entity.Post, tx *gorm.DB) error
	GetByID(id uint) (*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	Update(post *entity.Post, tx *gorm.DB) error
	Delete(id uint, tx *gorm.DB) error
}
