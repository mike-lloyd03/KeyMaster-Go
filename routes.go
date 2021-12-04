package main

import "keymaster_go/handlers"

func initializeRoutes() {
	router.GET("/", handlers.ShowIndex)
	router.GET("/keys", handlers.GetKeys)
	router.GET("/add_key", handlers.GetAddKey)
	router.POST("/add_key", handlers.PostAddKey)
	router.GET("/edit_key", handlers.GetEditKey)
	router.POST("/edit_key", handlers.PostEditKey)

	router.GET("/users", handlers.GetUsers)
	router.GET("/add_user", handlers.GetAddUser)
	router.POST("/add_user", handlers.PostAddUser)
	router.GET("/edit_user", handlers.GetEditUser)
	router.POST("/edit_user", handlers.PostEditUser)

	router.GET("/assignments", handlers.ShowAssignments)
	router.GET("/assign_key", handlers.ShowAssignKey)
}
