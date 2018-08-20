package pages

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.String(200, "Hello World")
}

func Licences(c *gin.Context) {
	c.String(200, "Licences")
}

func LicenceSpec(c *gin.Context) {
	c.String(200, "Licence %s", c.Param("licenceType"))
}

func ExamGen(c *gin.Context) {
	c.String(200, "Licences %s exam", c.Param("licenceType"))
}
