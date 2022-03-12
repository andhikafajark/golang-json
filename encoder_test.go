package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Master",
		MiddleName: "Zero",
		LastName:   "One",
	}

	_ = encoder.Encode(customer)
}
