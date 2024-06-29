package api_test

import (
	"fem/go/crypto/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error not found")
	}
}
