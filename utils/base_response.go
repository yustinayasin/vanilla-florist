package utils

import (
	"net/http"
)

// ReturnJsonResponse function for returning data in JSON format, bisa juga bikin costum format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json") //Header should be called first before WriteHeader and Write
	res.WriteHeader(httpCode)                            //mainly used to pass error code bcs the default one is 200
	res.Write(resMessage)
}
