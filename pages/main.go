package pages

import "github.com/gin-gonic/gin"

var (
	licenceCodeToName = map[string]string{
		"F":  "Foundation",
		"I":  "Intermediate",
		"AV": "Advanced",
	}
)

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
	c.HTML(200, "contact.html", nil)
}
func Faq(c *gin.Context) {
	c.HTML(200, "faq.html", nil)
}
func Robots(c *gin.Context) {
	c.HTML(200, "robots.txt", nil)
}
func Privacy(c *gin.Context) {
	c.HTML(200, "privacy.html", nil)
}
