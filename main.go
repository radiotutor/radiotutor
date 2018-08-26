package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := routes()
	r.LoadHTMLGlob("templates/*")

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8080")
}
