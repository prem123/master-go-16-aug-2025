package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

// Read JSON request body into a Go struct
func readRequestJSON(req *http.Request, target any) error {
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expect application/json Content-Type, got %s", mediaType)
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}

// Render response as JSON
func renderJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
