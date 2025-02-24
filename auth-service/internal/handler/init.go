package handler

import "net/http"

func InitServer() {
	http.HandleFunc("POST /register")
	http.HandleFunc("POST /login")
	http.HandleFunc("GET /validate")
}
