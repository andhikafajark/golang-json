package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Test","price":2000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func TestMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"Id":    "P0001",
		"Name":  "Test",
		"price": 2000000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
