package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/R-Media-Solutions/rmediasolutions-website/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)

	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
