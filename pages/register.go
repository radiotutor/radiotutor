package pages

import (
	"github.com/abaft/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pe5er/radiotutor/user"
)

func RegisterGET(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get("loggedIn")

	if v == nil {
		c.HTML(200, "register.html", nil)
	} else {
		c.HTML(200, "login-successful.html", gin.H{"User": v.(user.User)})
	}
}

func RegisterPOST(c *gin.Context) {
	session := sessions.Default(c)

	rawUsername, ok := c.GetPostForm("username")
	if !ok || rawUsername == "" {
		c.HTML(200, "register.html", gin.H{
			"ErrorTitle":   "ERROR",
			"ErrorMessage": "Need Username",
		})
		RegisterGET(c)
		return
	}
	rawEmail, ok := c.GetPostForm("email")
	if !ok || rawEmail == "" {
		// TODO Proper Email Validation
		c.HTML(200, "register.html", gin.H{
			"ErrorTitle":   "ERROR",
			"ErrorMessage": "Need Valid Email",
		})
		RegisterGET(c)
		return
	}
	rawPassword, ok := c.GetPostForm("password")
	if !ok || len(rawPassword) <= 5 {
		c.HTML(200, "register.html", gin.H{
			"ErrorTitle":   "ERROR",
			"ErrorMessage": "Need Password greater than 5 chars",
		})
		RegisterGET(c)
		return
	}

	u, err := user.CreateUser(rawUsername, rawPassword, rawEmail)
	if err != nil {
		c.HTML(200, "register.html", gin.H{
			"ErrorTitle":   "ERROR",
			"ErrorMessage": err.Error(),
		})
		return
	}
	session.Set("loggedIn", u)
	session.Save()
	RegisterGET(c)
}
