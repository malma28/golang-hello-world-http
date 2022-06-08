package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	request := httptest.NewRequest("GET", "/", nil)

	recorder := httptest.NewRecorder()
	index(recorder, request)

	result := recorder.Result()
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Hello World" {
		t.Fatal("Response is not \"Hello World\"")
	}
}
