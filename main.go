package main

import "todolist-api/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
