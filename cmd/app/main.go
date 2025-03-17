package main

import (
	"dynamo_db_crud/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
