package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// ReturnJsonResponse function for returning data in JSON format, bisa juga bikin costum format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json") //Header should be called first before WriteHeader and Write
	res.WriteHeader(httpCode)                            //mainly used to pass error code bcs the default one is 200
	res.Write(resMessage)
}

func ReturnErrorResponse(res http.ResponseWriter, httpCode int, errorMessage string) {
	// Create ErrorResponse struct instance
	errorResponse := ErrorResponse{
		Message: errorMessage,
	}

	// Convert ErrorResponse struct to JSON
	errorJSON, err := json.Marshal(errorResponse)
	if err != nil {
		// If error occurs during JSON marshaling, return internal server error
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	// Set response header and write response
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(errorJSON)
}
