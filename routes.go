package main

import (
	"encoding/gob"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/abaft/sessions"
	"github.com/abaft/sessions/redis"
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

	cookiesSessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	// Static Pages
	e.GET("/", cache.CachePage(cacheStore, time.Hour, p.Home))
	e.GET("/faq", cache.CachePage(cacheStore, time.Hour, p.Faq))
	e.GET("/contact", cache.CachePage(cacheStore, time.Hour, p.Contact))
	e.GET("/robots.txt", cache.CachePage(cacheStore, time.Hour, p.Robots))
	e.GET("/privacy", cache.CachePage(cacheStore, time.Hour, p.Privacy))

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

func HttpsRedirect() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		suffix := "/"
		for _, p := range c.Params {
			suffix += p.Value + "/"
		}
		c.Redirect(302, "https://radiotutor.uk"+suffix)
	})

	return e
}
