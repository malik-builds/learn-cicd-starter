package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T){
	// 1. Set up fake HTTP headers
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	//2 call the function
	key, err := GetAPIKey(headers)

	// 3 check the result
	if err != nil{
		t.Errorf("expected no error, got: %v", err)
	}
	if key!= "my-secret-key"{
		t.Errorf("expected 'my-secret-key', got: %q", key)
	}
}