package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHelloUser(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "Without name",
			query:    "",
			expected: "Hello User!",
		},
		{
			name:     "With name",
			query:    "?name=Malma",
			expected: "Hello Malma!",
		},
		{
			name:     "With invalid name 1",
			query:    "?name",
			expected: "Hello User!",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/"+testCase.query, nil)

			recorder := httptest.NewRecorder()
			index(recorder, request)

			result := recorder.Result()
			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(body) != testCase.expected {
				t.Fatalf("Expected \"%v\" but got \"%v\" instead", testCase.expected, string(body))
			}
		})
	}
}
