package service

import "golang_fiber/post/entity"

type PostService interface {
	Create(post *entity.Post) error          // 게시글 생성
	Read(id uint) (*entity.Post, error)      // 특정 게시글 조회
	List() ([]*entity.Post, error)           // 모든 게시글 조회
	Update(post *entity.Post) error         // 게시글 수정
	Delete(id uint) error                    // 게시글 삭제
}
