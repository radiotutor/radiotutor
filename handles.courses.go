// handlers.courses.go

package main

import (
	"github.com/gin-gonic/gin"
)

func showFoundationCourse(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Foundation"}, "foundation.html")
}