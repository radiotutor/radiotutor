package main

import (
	"encoding/gob"
	"github.com/abaft/sessions"
	"github.com/abaft/sessions/redis"
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/radiotutor/pages"
	"github.com/pe5er/radiotutor/quiz"
	"github.com/pe5er/radiotutor/user"
)

func routes() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/")
	})

	quiz.QuestionsInit()

	gob.Register([]quiz.Question{})
	gob.Register(user.User{})

	sessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	e.Use(sessions.Sessions("radioTutor", sessionStore))

	// Static Pages
	e.GET("/", p.Home)
	e.GET("/faq", p.Faq)
	e.GET("/contact", p.Contact)
	e.GET("/robots.txt", p.Robots)
	e.GET("/privacy", p.Privacy)
	e.GET("/news", p.News)
	e.GET("/donate", p.Donate)
	e.GET("/fcourse", p.FCourse)

	// user pages
	e.GET("/login", p.LoginGET)
	e.POST("/login", p.LoginPOST)
	e.GET("/register", p.RegisterGET)
	e.POST("/register", p.RegisterPOST)
	e.Any("/logout", p.Logout)
	// e.Any("/removeuser", p.RemoveUser) Removed due to concer over open sessions

	user := e.Group("/u/:username")
	{
		user.GET("page", p.AccountGET)
	}

	// Licences
	e.GET("/l", p.Licences)

	// Denominations
	licence := e.Group("/l/:licenceType", p.LicenceSpec)
	{
		licence.GET("exam", p.QuizGet)
		licence.POST("exam", p.QuizPost)
		licence.GET("course", p.Courses)
	}

	// Resource loading
	e.Static("/resources", "./resources")

	return e
}

func HttpsRedirect() *gin.Engine {
	e := gin.Default()
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "https://radiotutor.uk"+c.Request.URL.Path)
	})

	return e
}
