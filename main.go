package main

import (
	"fmt"
)

func main() {
	r := routes()
	r.LoadHTMLGlob("templates/*")

	quiz.QuestionsInit()

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8080")
}
