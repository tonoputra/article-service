package helper

import (
	"net/http"
	"strings"
)

type HTTPCodeInterface interface {
	MapError(e error) (int, string)
}

type httpCode struct {
}

func HTTPCode() HTTPCodeInterface {
	return &httpCode{}
}

// MapError return http status code based on error
func (httpCode) MapError(e error) (int, string) {
	switch {
	case strings.Contains(e.Error(), "context deadline exceeded"):
		return http.StatusRequestTimeout, "Request Server Timeout"
	case strings.Contains(e.Error(), "mongo: no documents in result"):
		return http.StatusBadRequest, "Document not found"
	default:
		return http.StatusInternalServerError, "Internal Server Error"
	}
}
