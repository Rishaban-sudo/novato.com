package json

import (
	"encoding/json"
	"fmt"
)

type JsonObject struct {
	data map[string]interface{}
}

func NewJsonObject() *JsonObject {
	return &JsonObject{
		data: make(map[string]interface{}),
	}
}

func (o *JsonObject) AddOrUpdate(key string, value interface{}) {
	o.data[key] = value
}

func (o *JsonObject) Get(key string) interface{} {
	return o.data[key]
}

func (o *JsonObject) Remove(key string) interface{} {
	deletedValue := o.data[key]
	delete(o.data, key)
	return deletedValue
}

func (o *JsonObject) GetJsonRepresentation() string {

	if o == nil || o.data == nil || len(o.data) == 0 {
		return "{}"
	}

	jsonString, err := json.MarshalIndent(o.data, "", " ")

	if err != nil {
		fmt.Println("[JsonObject] Failed to marshal !!")
		fmt.Println(err)
		return "{}"
	}

	return string(jsonString)
}
