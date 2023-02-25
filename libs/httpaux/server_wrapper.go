package httpaux

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"route256/libs/errors"
	"route256/libs/validator"
)

var (
	validationErr = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "invalid request data",
	}

	decodingErr = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "json decoding error",
	}

	encodingErr = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "json decoding error",
	}
)

type Wrapper[Req, Resp any] struct {
	fn func(ctx context.Context, req Req) (Resp, error)
}

func New[Req, Resp any](fn func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}

func (w *Wrapper[Req, Res]) ServeHTTP(resWriter http.ResponseWriter, httpReq *http.Request) {
	ctx := httpReq.Context()

	limitedReader := io.LimitReader(httpReq.Body, 1_000_000)

	var request Req
	err := json.NewDecoder(limitedReader).Decode(&request)
	if err != nil {
		errors.WriteHttpError(decodingErr, resWriter)
		return
	}

	if validator.Validate(request) {
		errors.WriteHttpError(validationErr, resWriter)
		return
	}

	response, err := w.fn(ctx, request)
	if err != nil {
		log.Printf("%s: error: %s\n", httpReq.URL, err)
		errors.WriteHttpError(err, resWriter)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		errors.WriteHttpError(encodingErr, resWriter)
		return
	}

	resWriter.Header().Add("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	_, _ = resWriter.Write(rawJSON)
}
