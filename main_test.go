package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	rec := httptest.NewRecorder()

	sleep()(rec, req)

	res := rec.Result()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	if !strings.Contains(string(b), "Slept") {
		t.Fatal("\"Slept\" missing")
	}
}
