package json_example

import (
	"testing"
)

func TestUnmarshalByInterface(t *testing.T) {
	var jsonBlob = []byte(`[
		{"ProductName": "Mac", "Amount": 1, "Info": {"Category": "Laptop"}},
		{"ProductName": "iPhone",    "Amount": 2, "Info": {"Category": "SmartPhone"}}
	]`)
	o, err := UnmarshalByInterface(jsonBlob)
	if err != nil {
		t.Error("UnmarshalByInterface(): unmarshal failed")
	}
	p := o.([]interface{})[0].(map[string]interface{})["ProductName"].(string)
	if p != "Mac" {
		t.Error("UnmarshalByInterface(): ProductName failed")
	}
	a := o.([]interface{})[1].(map[string]interface{})["Amount"].(float64)
	if a != 2 {
		t.Error("UnmarshalByInterface(): ProductName failed")
	}
	c := o.([]interface{})[0].(map[string]interface{})["Info"].(map[string]interface{})["Category"].(string)
	if c != "Laptop" {
		t.Error("UnmarshalByInterface(): Category failed")
	}
}
