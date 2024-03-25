package main

import (
	"fmt"
	"github.com/igorXimeness/educ-off-api/internal/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Server listerning on htpp://localhost%s\n", port)

	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
