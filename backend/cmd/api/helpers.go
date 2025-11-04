package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) writeJson(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func generateShortCode(longURL string) string {
	hash := 0
	for _, char := range longURL {
		hash = (hash << 5) - hash + int(char)
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte
	if hash == 0 {
		return string(charset[0])
	}

	hash = abs(hash)
	for hash > 0 {
		result = append(result, charset[hash%62])
		hash /= 62
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
