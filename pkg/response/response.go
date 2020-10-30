package response

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Map map[string]interface{}

func EJSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	if data == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		//return nil
	}

	j, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(j)
	//return nil
}

func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {

	msg := ErrorMessage{
		Message: message,
	}

	EJSON(w, r, statusCode, msg)
}
