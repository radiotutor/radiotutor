package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := routes()
	r.LoadHTMLGlob("templates/*")

	// Logging to a file.
	// Listen and Server in https://127.0.0.1:8080
	if os.Getenv("RADIOTUTOR_DEV") == "true" {
		fmt.Printf("hello world")
		r.Run(":8080")
	} else {

		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)

		autotls.Run(r, "radiotutor.uk")
	}
}
