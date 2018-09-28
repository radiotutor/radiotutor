package pages

import (
	"github.com/abaft/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pe5er/radiotutor/user"
)

var (
	licenceCodeToName = map[string]string{
		"F":  "Foundation",
		"I":  "Intermediate",
		"AV": "Advanced",
	}
)

func sHTML(c *gin.Context, code int, name string, obj gin.H) {
	session := sessions.Default(c)

	v := session.Get("loggedIn")
	if v != nil {
		uD := gin.H{"User": v.(user.User)}
		if obj == nil {
			c.HTML(code, name, uD)
			return
		}
		for k, v := range uD {
			obj[k] = v
		}
	}
	c.HTML(code, name, obj)
}

func Licences(c *gin.Context) {
	c.String(200, "Licences")
}

func LicenceSpec(c *gin.Context) {
	if s := c.Param("licenceType"); s != "F" && s != "I" && s != "AV" {
		c.Redirect(302, "/")
	}
}

func ExamGen(c *gin.Context) {
	c.String(200, "Licences %s exam", c.Param("licenceType"))
}

func Contact(c *gin.Context) {
	sHTML(c, 200, "contact.html", nil)
}
func Faq(c *gin.Context) {
	sHTML(c, 200, "faq.html", nil)
}
func Robots(c *gin.Context) {
	sHTML(c, 200, "robots.txt", nil)
}
func Privacy(c *gin.Context) {
	sHTML(c, 200, "privacy.html", nil)
}
func News(c *gin.Context) {
	sHTML(c, 200, "newsstatic.html", nil)
}
func Donate(c *gin.Context) {
	sHTML(c, 200, "donate.html", nil)
}

