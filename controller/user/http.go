package handler

import (
 "native-go-api/models"
 "native-go-api/db"
 "native-go-api/utils"
 "net/http"
 "encoding/json"
)

// Add a movie handler
func AddMovie(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)
	
		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}
   
	var movie models.Movie
   
	payload := req.Body
   
	defer req.Body.Close()
	// parse the movie data into json format
	err := json.NewDecoder(payload).Decode(&movie)
	
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the movie data",
		}`)
	
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}
   
	db.Moviedb[movie.ID] = movie
	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "Movie was successfully created",
	 }`)
   
	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
   }