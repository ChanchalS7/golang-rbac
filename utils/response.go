package utils

import (
	"net/http"
	"encoding/json"
)

//RespondWithError send an error response
func ResponseWithError(w http.ResponseWriter, code int , message string){
	RespondWithJSON(w, code, map[string]string{"error":message})
}
//RespondWitJSON send a JSON response
func RespondWithJSON( w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(code)
	w.Write(response)
}