package repository

import (
	"golang_fiber/post/entity"
	"gorm.io/gorm"
)

// PostRepositoryImpl은 PostRepository 인터페이스를 구현하는 구조체
type PostRepositoryImpl struct {
	DB *gorm.DB
}

// NewPostRepositoryImpl은 PostRepositoryImpl을 초기화하는 함수
func NewPostRepositoryImpl(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{DB: db}
}

// Create 게시글을 DB에 저장
func (repository *PostRepositoryImpl) Create(post *entity.Post, tx *gorm.DB) error {
	return tx.Create(post).Error
}

// GetByID 특정 ID의 게시글 조회
func (repository *PostRepositoryImpl) GetByID(id uint) (*entity.Post, error) {
	var post entity.Post
	result := repository.DB.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

// GetAll 모든 게시글 조회
func (repository *PostRepositoryImpl) GetAll() ([]*entity.Post, error) {  // 수정된 부분
	var posts []*entity.Post  // 수정된 부분
	result := repository.DB.Find(&posts)
	return posts, result.Error
}

// Update 게시글 수정
func (repository *PostRepositoryImpl) Update(post *entity.Post, tx *gorm.DB) error {
	return tx.Save(post).Error
}

// Delete 게시글 삭제
func (repository *PostRepositoryImpl) Delete(id uint, tx *gorm.DB) error {
	return tx.Delete(&entity.Post{}, id).Error
}
