package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/radiotutor/pages"
	"time"
)

func routes() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	store := persistence.NewInMemoryStore(time.Second)

	// Homepage
	e.GET("/", cache.CachePage(store, time.Hour, p.Home))

	// Licences
	e.GET("/l", p.Licences)

	// Denominations
	licence := e.Group("/l/:licenceType", p.LicenceSpec)
	{
		licence.GET("exam", cache.CachePage(store, time.Hour, p.ExamGen))
		licence.GET("course", cache.CachePage(store, time.Hour, p.Courses))
	}

	// Resource loading
	e.Static("/resources", "./resources")

	return e
}
