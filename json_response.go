package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleSuccessJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("%v occured while sending response %v", err, payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func HandleErrorJson(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Internal Error %v", msg)
	}

	type errRes struct {
		Error string `json:"error"`
	}

	HandleSuccessJson(w, code, errRes{Error: msg})

}
