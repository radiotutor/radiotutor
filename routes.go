package main

import (
	"encoding/gob"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/radiotutor/pages"
	"github.com/pe5er/radiotutor/quiz"
	"time"
)

func routes() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	quiz.QuestionsInit()

	cacheStore := persistence.NewInMemoryStore(time.Second)
	gob.Register([]quiz.Question{})

	cookiesSessionStore := memstore.NewStore([]byte("HANGE_IN_PRODUCTION"))

	// Static Pages
	e.GET("/", cache.CachePage(cacheStore, time.Hour, p.Home))
	e.GET("/faq", cache.CachePage(cacheStore, time.Hour, p.Faq))
	e.GET("/contact", cache.CachePage(cacheStore, time.Hour, p.Contact))

	// Licences
	e.GET("/l", p.Licences)

	// Denominations
	licence := e.Group("/l/:licenceType", p.LicenceSpec)
	{
		licence.GET("exam", sessions.Sessions("quiz", cookiesSessionStore), p.QuizGet)
		licence.POST("exam", sessions.Sessions("quiz", cookiesSessionStore), p.QuizPost)
		licence.GET("course", cache.CachePage(cacheStore, time.Hour, p.Courses))
	}

	// Resource loading
	e.Static("/resources", "./resources")

	return e
}
