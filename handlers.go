package main

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	HandleSuccessJson(w, 200, struct{}{})
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	HandleErrorJson(w, 400, "Something went wrong")
	
}