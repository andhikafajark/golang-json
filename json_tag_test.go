package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Test",
		ImageURL: "http://test.com/example.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Test","image_url":"http://test.com/example.png"}`
	jsonBytes := []byte(jsonString)

	customer := &Product{}

	_ = json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}
