package pages

func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
