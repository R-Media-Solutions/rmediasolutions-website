package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
