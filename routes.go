package main

import (
	p "github.com/pe5er/radiotutor/pages"
	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	// Homepage
	e.GET("/", p.Home)

	// Licences
	e.GET("/licence", p.Licences)

	// Denominations
	licence := e.Group("/licence/:licenceType", p.LicenceSpec)
	{
		licence.GET("exam", p.ExamGen)
	}

	// Resource loading
	e.Static("/resources", "./resources")
	return e
}
