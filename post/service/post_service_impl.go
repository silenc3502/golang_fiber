package service

import (
	"golang_fiber/post/entity"
	"golang_fiber/post/repository"
	"gorm.io/gorm"
)

// PostService는 비즈니스 로직을 처리
type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *gorm.DB
}

// NewPostService 생성자 함수
func NewPostService(postRepo repository.PostRepository, db *gorm.DB) *PostServiceImpl {
	return &PostServiceImpl{PostRepository: postRepo, DB: db}
}

// NewPostService 생성자 함수
func (service *PostServiceImpl) Create(post *entity.Post) error {
	tx := service.DB.Begin() // 트랜잭션 시작
	if tx.Error != nil {
		return tx.Error
	}

	// 트랜잭션 내에서 데이터베이스 작업 수행
	if err := service.PostRepository.Create(post, tx); err != nil {
		tx.Rollback() // 오류 발생 시 롤백
		return err
	}

	tx.Commit() // 트랜잭션 커밋
	return nil
}

// Read 특정 게시글 조회
func (service *PostServiceImpl) Read(id uint) (*entity.Post, error) {
	return service.PostRepository.GetByID(id)
}

// List 모든 게시글 조회
func (service *PostServiceImpl) List() ([]*entity.Post, error) {
	return service.PostRepository.GetAll()
}

// Update 게시글 수정
func (service *PostServiceImpl) Update(post *entity.Post) error {
	tx := service.DB.Begin() // 트랜잭션 시작
	if tx.Error != nil {
		return tx.Error
	}

	// 트랜잭션 내에서 데이터베이스 작업 수행
	if err := service.PostRepository.Update(post, tx); err != nil {
		tx.Rollback() // 오류 발생 시 롤백
		return err
	}

	tx.Commit() // 트랜잭션 커밋
	return nil
}

// Delete 게시글 삭제
func (service *PostServiceImpl) Delete(id uint) error {
	tx := service.DB.Begin() // 트랜잭션 시작
	if tx.Error != nil {
		return tx.Error
	}

	// 트랜잭션 내에서 데이터베이스 작업 수행
	if err := service.PostRepository.Delete(id, tx); err != nil {
		tx.Rollback() // 오류 발생 시 롤백
		return err
	}

	tx.Commit() // 트랜잭션 커밋
	return nil
}
