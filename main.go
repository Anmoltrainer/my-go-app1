package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		htmlContent := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Colorful Sample Website</title>
			</head>
			<body>
				<h1 style="color: blue;">Welcome to the Colorful Sample Website for sudhan!</h1>
				<p style="color: green;">This is a simple example of a colorful website served by a Go web application.</p>
				<p style="color: red;">Enjoy the colors!</p>
			</body>
			</html>
		`
		fmt.Fprint(w, htmlContent)
	}).Methods("GET")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
