package pages

import "github.com/gin-gonic/gin"

func Courses(c *gin.Context) {

	switch c.Param("licenceType") {
	case "M0":
		c.HTML(200, "advanced.html", nil)
	case "2E0":
		c.HTML(200, "intermediate.html", nil)
	case "M6":
		c.HTML(200, "foundation.html", nil)
	}
}
