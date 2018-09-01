package pages

import "github.com/gin-gonic/gin"

func Courses(c *gin.Context) {

	switch c.Param("licenceType") {
	case "AV":
		c.HTML(200, "advanced.html", nil)
	case "I":
		c.HTML(200, "intermediate.html", nil)
	case "F":
		c.HTML(200, "foundation.html", nil)
	}
}
