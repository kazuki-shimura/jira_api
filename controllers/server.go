package controllers

import (
	"fmt"
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "これはTopページです")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "これはAboutページです")
}

func StartServer() error {
	http.HandleFunc("/", Top)
	http.HandleFunc("/about", About)
	return http.ListenAndServe(":5000", nil)
}
