package main

func main() {
	r := routes()

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8080")
}
