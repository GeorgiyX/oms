package errors

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type HttpError struct {
	Code    int
	Message string `Message:"error"`
}

var (
	errFallBack = HttpError{
		Code:    http.StatusInternalServerError,
		Message: "internal error",
	}
	errInternalBytes = []byte("{\n\t\"error\": \"internal error\"\n}")
)

func (e HttpError) Error() string {
	return e.Message
}

func WriteHttpError(errIn error, w http.ResponseWriter) {
	var httpErr HttpError
	if !errors.As(errIn, &httpErr) {
		httpErr = errFallBack
	}

	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(httpErr)
	if err != nil {
		log.Printf("error marshal err: %v\n", errIn)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errInternalBytes)
		return
	}

	w.WriteHeader(httpErr.Code)
	w.Write(bytes)
}
