package pages

import "github.com/gin-gonic/gin"

func Courses(c *gin.Context) {

	switch c.Param("licenceType") {
	case "AV":
		sHTML(c, 200, "advanced.html", nil)
	case "I":
		sHTML(c, 200, "intermediate.html", nil)
	case "F":
		sHTML(c, 200, "foundation.html", nil)
	}
}
