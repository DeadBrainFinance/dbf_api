package utils

import (
	// "encoding/json"
	"net/http"
)

type Response struct {
    StatusCode int
    Msg string
}

func WriteResponse(w http.ResponseWriter, statusCode int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if msg != "" {
        w.Write([]byte(msg))
    }
    // response := Response{
    //     StatusCode: statusCode,
    //     Msg: msg,
    // }
    // json.NewEncoder(w).Encode(response)
}
