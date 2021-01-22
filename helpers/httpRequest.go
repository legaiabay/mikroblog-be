package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

const HTTPStatusSuccess string = "success"
const HTTPStatusError string = "error"

//HTTPResponse - set response
func HTTPResponse(w http.ResponseWriter, status string, statusCode int, message string) {
	if status == HTTPStatusSuccess {
		Logger.Info(message)
	} else {
		Logger.Error(message)
	}

	d := Response{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(d)

	return
}

//ToJSON -> convert struct to JSON
func ToJSON(d interface{}) string {
	j, _ := json.Marshal(d)
	return string(j)
}
