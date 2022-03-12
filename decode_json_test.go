package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Master","MiddleName":"Zero","LastName":"One"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	_ = json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}
