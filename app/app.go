package app

import (
	"../domain"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnDB()
	router.Run(":8080")
}
