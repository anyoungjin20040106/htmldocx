package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 포트 설정
	port := fmt.Sprintf(":%d", 1234)

	// Gin 라우터 생성
	router := gin.Default()

	// 정적 파일 서버 설정 ("/chapter" 디렉토리를 "/c" 경로로 제공)
	router.Static("/c", "./chapter")

	// 기본 라우트 설정 (index.html 제공)
	router.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	// /doc 라우트 설정 (POST 요청 처리)
	router.POST("/doc", func(c *gin.Context) {
		pw := c.PostForm("pw")
		if pw != os.Getenv("pw") {
			c.String(http.StatusUnauthorized, "암호가 틀렸습니다")
		} else {
			c.File("doc.html")
		}
	})

	// 서버 시작
	log.Println("Server started on " + port)
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
