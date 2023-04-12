package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"path/filepath"
)

func GetJsonBody(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Invalid Content-Type")
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return nil, err
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)

	if err != nil && err != io.EOF {
		return nil, err
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		return nil, err
	}

	if len(jsonBody) == 0 {
		return nil, errors.New("Empty request body")
	}

	return jsonBody, nil
}

func GetURLID(r *http.Request, url string) string {
	sub := strings.TrimPrefix(r.URL.Path, url)
	_, id := filepath.Split(sub)
	return id
}