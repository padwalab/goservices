package responsehandler

import (
	"encoding/json"
	"net/http"

	"github.com/padwalab/goservices/logger"
)

// WriteResponse writes the response
func WriteResponse(w http.ResponseWriter, rl ResponseLevel, data interface{}, message string) {
	switch rl {
	case ERROR:
		errorResponse(w, data, message)
		break
	case OK:
		successResponse(w, data, message)
		break
	}
}

// SuccessResponse writes the success response
func successResponse(w http.ResponseWriter, data interface{}, message string) {
	setResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	res := Response{
		Status:  "SUCCESS",
		Data:    data,
		Message: message,
	}
	logger.Log("INFO", res)
	response, err := json.Marshal(res)
	if err != nil {
		response, _ := json.Marshal(Response{
			Status:  "error",
			Message: "Unable to marshal server response",
		})
		w.Write(response)
		return
	}
	w.Write(response)
	return
}

// ErrorResponse writes the error response
func errorResponse(w http.ResponseWriter, data interface{}, message string) {
	setResponseHeaders(w)
	w.WriteHeader(http.StatusBadRequest)
	res := Response{
		Status:  "ERROR",
		Data:    data,
		Message: message,
	}
	logger.Log(logger.ERROR, res)
	response, err := json.Marshal(res)
	if err != nil {
		response, _ := json.Marshal(Response{
			Status:  "error",
			Message: "Unable to marshal server response",
		})
		w.Write(response)
		return
	}
	w.Write(response)
	return
}

// SetResponseHeaders ...
func setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
