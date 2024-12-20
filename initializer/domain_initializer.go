package initializer

import (
	"fmt"
	"golang_fiber/post/controller"
	"golang_fiber/post/entity"
	"golang_fiber/post/repository"
	"golang_fiber/post/service"
	"os"

	"github.com/google/wire"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var PostSet = wire.NewSet(
	repository.NewPostRepositoryImpl,
	service.NewPostService,
	controller.NewPostController,
)

// DomainInitializer는 데이터베이스와 서비스 객체를 초기화하는 함수입니다.
func DomainInitializer() (*gorm.DB, service.PostService, *controller.PostController, error) {
	// .env 파일 로딩
	if err := godotenv.Load(); err != nil {
		return nil, nil, nil, fmt.Errorf("Error loading .env file")
	}

	// .env에서 MySQL 연결 정보를 가져옵니다.
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbLoc := os.Getenv("DB_LOC")

	// DSN (Data Source Name) 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbLoc)

	// MySQL 데이터베이스 연결
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	// 테이블 자동 마이그레이션
	if err := db.AutoMigrate(&entity.Post{}); err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to auto-migrate: %v", err)
	}

	wire.Build(PostSet)
	// return nil, nil
}
