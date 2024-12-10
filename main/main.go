package main

import (
	"fmt"
	"golang_fiber/initializer"
)

func main() {
	_, _, postController, err := initializer.DomainInitializer()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fiber 앱 초기화 및 라우팅 설정
	app := initializer.AppInitializer(postController)

	// 서버 실행
	app.Listen(":3773")
}
