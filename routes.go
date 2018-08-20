// routes.go

package main

func initializeRoutes() {

	router.Use(setUserStatus())
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{

		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	courseRoutes := router.Group("/c")
	{
		courseRoutes.GET("/foundation", showFoundationCourse)
		courseRoutes.GET("/intermediate", showIntermediateCourse)
		courseRoutes.GET("/advanced", showAdvancedCourse)
	}
}
