package app

import (
	"../controllers/ping"
	"../controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/find", users.FindUser)
	router.GET("/users/delete", users.DeleteUser)
	router.GET("/users/update", users.UpdateUser)
	router.POST("/users/create", users.CreateUser)
	router.POST("/linkData/create", users.CreateLinkData)
}
