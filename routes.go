package main

func initializeRoutes() {
	router.GET("/", showIndex)
	router.GET("/keys", showKeys)
	router.GET("/users", showUsers)
}
