package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	requestBody := `{
		"k1": 7,
		"k2": ["v2", "v3", {
			"k4": "v4",
			"k5": ["v5", "v6"]
		}],
		"k3": {
			"k4": "v4",
			"k5": ["v5", "v6"]
		}
	}`

	var body map[string]interface{}
	err := json.Unmarshal([]byte(requestBody), &body)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	checkValue(body)

	res, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(res))
}

func checkValue(requestBody interface{}) {
	switch val := requestBody.(type) {
	case map[string]interface{}:
		for k, v := range val {
			checkValue(v)
			if str, ok := v.(string); ok {
				val[k] = reverseString(str)
			}
		}
	case []interface{}:
		for i, v := range val {
			checkValue(v)
			if str, ok := v.(string); ok {
				val[i] = reverseString(str)
			}
		}
	}
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
