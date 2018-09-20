package pages

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	sHTML(c, 200, "index.html", nil)
}
