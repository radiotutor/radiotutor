package main

func main() {
	r := routes()
	r.LoadHTMLGlob("templates/*")

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8080")
}
