package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	handler = handlerCreateEvent{&service{}}
)

func TestHandlerCreateEvent_ServeHTTP(t *testing.T) {

	value := url.Values{}
	value.Add("user_id", "3")
	value.Add("name", "go school")
	value.Add("data", "2019-09-09")

	bodyRequest := strings.NewReader(value.Encode())

	req := httptest.NewRequest("POST", "/create_event", bodyRequest)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидалось: %d\n результат: %d", http.StatusOK, resp.StatusCode)
	}
}
