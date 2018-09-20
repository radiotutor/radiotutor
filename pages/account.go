package pages

import (
	//"github.com/abaft/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/pe5er/radiotutor/user"
)

func AccountGET(c *gin.Context) {
	sHTML(c, 200, "account.html", nil)
}
