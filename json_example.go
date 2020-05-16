package json_example

import (
	"encoding/json"
	"fmt"
)

func Unmarshal(b []byte) (interface{}, error) {
	var orders interface{}
	err := json.Unmarshal(b, &orders)
	if err != nil {
		fmt.Println("json unmarshal was failed")
	}
	return orders, err
}
