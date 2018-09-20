package main

import (
	"encoding/gob"
	"github.com/abaft/sessions"
	"github.com/abaft/sessions/redis"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/radiotutor/pages"
	"github.com/pe5er/radiotutor/quiz"
	"github.com/pe5er/radiotutor/user"
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
	gob.Register(user.User{})

	sessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	e.Use(sessions.Sessions("radioTutor", sessionStore))

	// Static Pages
	e.GET("/", cache.CachePage(cacheStore, time.Hour, p.Home))
	e.GET("/faq", cache.CachePage(cacheStore, time.Hour, p.Faq))
	e.GET("/contact", cache.CachePage(cacheStore, time.Hour, p.Contact))
	e.GET("/robots.txt", cache.CachePage(cacheStore, time.Hour, p.Robots))
	e.GET("/privacy", cache.CachePage(cacheStore, time.Hour, p.Privacy))

	// user pages
	e.GET("/login", p.LoginGET)
	e.POST("/login", p.LoginPOST)
	e.GET("/register", p.RegisterGET)
	e.POST("/register", p.RegisterPOST)
	e.Any("/logout", p.Logout)
	e.Any("/removeuser", p.RemoveUser)

	user := e.Group("/u/:username")
	{
		user.GET("page", p.UserPage)
	}

	// Licences
	e.GET("/l", p.Licences)

	// Denominations
	licence := e.Group("/l/:licenceType", p.LicenceSpec)
	{
		licence.GET("exam", p.QuizGet)
		licence.POST("exam", p.QuizPost)
		licence.GET("course", cache.CachePage(cacheStore, time.Hour, p.Courses))
	}

	// Resource loading
	e.Static("/resources", "./resources")

	return e
}

func HttpsRedirect() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "https://radiotutor.uk"+c.Request.URL.Path)
	})

	return e
}
