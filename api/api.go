package api

import (
	"io"
	"net/http"
)

func FetchAPI(api string) ([]byte, error) {
	response, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
