package client

import (
	"io"
	"net/http"
)

func execute() ([]byte, error) {
	response, err := http.Get("http://localhost:8080/ping")

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
