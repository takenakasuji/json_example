package json_example

import (
	"testing"

	"github.com/koron/go-dproxy"
	"github.com/mattn/go-jsonpointer"
)

func TestUnmarshalByInterface(t *testing.T) {
	var jsonBlob = []byte(`[
		{"ProductName": "Mac", "Amount": 1, "Info": {"Category": "Laptop", "Tags": ["New", "13inch"]}},
		{"ProductName": "iPhone", "Amount": 2, "Info": {"Category": "SmartPhone", "Tags": ["6inch", "New"]}}
	]`)
	o, err := Unmarshal(jsonBlob)
	if err != nil {
		t.Error("UnmarshalByInterface(): unmarshal failed")
	}
	p := o.([]interface{})[0].(map[string]interface{})["ProductName"].(string)
	if p != "Mac" {
		t.Error("UnmarshalByInterface(): ProductName failed")
	}
	a := o.([]interface{})[1].(map[string]interface{})["Amount"].(float64)
	if a != 2 {
		t.Error("UnmarshalByInterface(): Amount failed")
	}
	c := o.([]interface{})[0].(map[string]interface{})["Info"].(map[string]interface{})["Category"].(string)
	if c != "Laptop" {
		t.Error("UnmarshalByInterface(): Category failed")
	}
	tg := o.([]interface{})[0].(map[string]interface{})["Info"].(map[string]interface{})["Tags"].([]interface{})[0].(string)
	if tg != "New" {
		t.Error("UnmarshalByInterface(): Tags failed")
	}
}

func TestUnmarshalByDproxy(t *testing.T) {
	var jsonBlob = []byte(`[
		{"ProductName": "Mac", "Amount": 1, "Info": {"Category": "Laptop", "Tags": ["New", "13inch"]}},
		{"ProductName": "iPhone", "Amount": 2, "Info": {"Category": "SmartPhone", "Tags": ["6inch", "New"]}}
	]`)
	o, err := Unmarshal(jsonBlob)
	if err != nil {
		t.Error("TestUnmarshalByDproxy(): unmarshal failed")
	}
	p, _ := dproxy.New(o).A(0).M("ProductName").String()
	if p != "Mac" {
		t.Error("TestUnmarshalByDproxy(): ProductName failed")
	}
	a, _ := dproxy.New(o).A(1).M("Amount").Float64()
	if a != 2 {
		t.Error("TestUnmarshalByDproxy(): Amount failed")
	}
	c, _ := dproxy.New(o).A(0).M("Info").M("Category").String()
	if c != "Laptop" {
		t.Error("TestUnmarshalByDproxy(): Category failed")
	}
	tg, _ := dproxy.New(o).A(0).M("Info").M("Tags").A(0).String()
	if tg != "New" {
		t.Error("TestUnmarshalByDproxy(): Tags failed")
	}
}

func TestUnmarshalByJsonPointer(t *testing.T) {
	var jsonBlob = []byte(`[
		{"ProductName": "Mac", "Amount": 1, "Info": {"Category": "Laptop", "Tags": ["New", "13inch"]}},
		{"ProductName": "iPhone", "Amount": 2, "Info": {"Category": "SmartPhone", "Tags": ["6inch", "New"]}}
	]`)
	o, err := Unmarshal(jsonBlob)
	if err != nil {
		t.Error("TestUnmarshalByJsonPointer(): unmarshal failed")
	}
	p, _ := jsonpointer.Get(o, "/0/ProductName")
	if p != "Mac" {
		t.Error("TestUnmarshalByJsonPointer(): ProductName failed")
	}
	a, _ := jsonpointer.Get(o, "/1/Amount")
	if a.(float64) != 2 {
		t.Error("TestUnmarshalByJsonPointer(): Amount failed")
	}
	c, _ := jsonpointer.Get(o, "/0/Info/Category")
	if c != "Laptop" {
		t.Error("TestUnmarshalByJsonPointer(): Category failed")
	}
	tg, _ := jsonpointer.Get(o, "/0/Info/Tags/0")
	if tg != "New" {
		t.Error("TestUnmarshalByJsonPointer(): Tags failed")
	}
}
