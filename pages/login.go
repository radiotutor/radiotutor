package pages

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
