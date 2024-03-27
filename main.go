package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/update/", update)
	http.HandleFunc("/delete/", delete)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
