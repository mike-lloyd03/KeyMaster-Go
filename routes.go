package main

import "keymaster_go/handlers"

func initializeRoutes() {
	router.GET("/", handlers.ShowIndex)
	router.GET("/keys", handlers.ShowKeys)
	router.GET("/add_key", handlers.ShowAddKey)
	router.POST("/add_key", handlers.PostAddKey)
	router.GET("/users", handlers.ShowUsers)
	router.GET("/assignments", handlers.ShowAssignments)
	router.GET("/assign_key", handlers.ShowAssignKey)
}
