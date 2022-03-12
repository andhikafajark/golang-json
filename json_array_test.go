package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Master",
		MiddleName: "Zero",
		LastName:   "One",
		Hobbies:    []string{"Game", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Master","MiddleName":"Zero","LastName":"One","Hobbies":["Game","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	_ = json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Master",
		MiddleName: "Zero",
		LastName:   "One",
		Hobbies:    []string{"Game", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Unknown Street 1",
				Country:    "Unknown Country 1",
				PostalCode: "Unknown PostalCode 1",
			},
			{
				Street:     "Unknown Street 2",
				Country:    "Unknown Country 2",
				PostalCode: "Unknown PostalCode 2",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Master","MiddleName":"Zero","LastName":"One","Hobbies":["Game","Reading","Coding"],"Addresses":[{"Street":"Unknown Street 1","Country":"Unknown Country 1","PostalCode":"Unknown PostalCode 1"},{"Street":"Unknown Street 2","Country":"Unknown Country 2","PostalCode":"Unknown PostalCode 2"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	_ = json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Unknown Street 1","Country":"Unknown Country 1","PostalCode":"Unknown PostalCode 1"},{"Street":"Unknown Street 2","Country":"Unknown Country 2","PostalCode":"Unknown PostalCode 2"}]`
	jsonBytes := []byte(jsonString)

	customer := &[]Address{}

	_ = json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Unknown Street 1",
			Country:    "Unknown Country 1",
			PostalCode: "Unknown PostalCode 1",
		},
		{
			Street:     "Unknown Street 2",
			Country:    "Unknown Country 2",
			PostalCode: "Unknown PostalCode 2",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
