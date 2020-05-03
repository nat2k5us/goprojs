package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	req, err := http.NewRequest("Get", "localhost:8080/Hello", nil)
	if err != nil {
		t.Fatal("Failed http request")
	}

	rec := httptest.NewRecorder()
	GreetingHandler(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected Status OK but got : %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected Status OK but got : %v", res.StatusCode)
	}

	fmt.Printf(string(b))
}

func TestGreetingHandler(t *testing.T) {

	tests := []struct {
		name     string
		endpoint string
		result   string
	}{
		{name: "good case", endpoint: "Hello", result: "Hello from the other side!!!"},		
	}
	for _, ts := range tests {
		req, err := http.NewRequest("Get", "localhost:8080/" + ts.endpoint, nil)
		if err != nil {
			t.Fatal("Failed http request")
		}

		rec := httptest.NewRecorder()
		GreetingHandler(rec, req)

		res := rec.Result()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected Status OK but got : %v", res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected Status OK but got : %v", res.StatusCode)
		}

		fmt.Printf(string(b))
	}
}
