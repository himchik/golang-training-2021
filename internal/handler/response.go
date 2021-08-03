package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error":"Internal server error"}`))
		if err != nil {
			log.Println("error: http.ResponseWriter.Write() func:", err)
		}
		return
	}
	w.WriteHeader(code)
	_, err = w.Write([]byte(b))
	if err != nil {
		log.Println("error: http.ResponseWriter.Write() func:", err)
	}
}
