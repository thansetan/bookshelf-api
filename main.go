package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":9000"
	fmt.Printf("Server running at : localhost%s\n", port)
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookHandler)
	http.ListenAndServe(port, nil)
}
