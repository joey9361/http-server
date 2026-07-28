package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
)

func TestHandleHello(t *testing.T) {
	w := httptest.NewRecorder()

	handleHello(w, nil)

	desiredCode := http.StatusOK
	if w.Code == desiredCode {
		t.Errorf("bad response code, expected %v but got %v\nbody: %s\n", 
				desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello world!")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("Expected: %s\nGot: %s", expectedMessage, w.Body.String())
	}
}