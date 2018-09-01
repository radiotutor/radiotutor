package pages

import "github.com/gin-gonic/gin"

func Licences(c *gin.Context) {
	c.String(200, "Licences")
}

func LicenceSpec(c *gin.Context) {
	if s := c.Param("licenceType"); s != "M0" && s != "2E0" && s != "M6" {
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
