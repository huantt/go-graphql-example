package graphql

import (
	"encoding/json"
)

type Json map[string]interface{}

func (Json) ImplementsGraphQLType(name string) bool {
	return name == "Json"
}

func (m *Json) UnmarshalGraphQL(input interface{}) error {
	if input == nil {
		return nil
	}
	var bytes []byte
	switch inputType := input.(type) {
	case []byte:
		bytes = inputType
	default:
		var err error
		bytes, err = json.Marshal(input)
		if err != nil {
			return err
		}
	}
	if string(bytes) == "\"\"" {
		return nil
	}
	return json.Unmarshal(bytes, m)
}

func NewJson(input interface{}) (*Json, error) {
	if input == nil {
		return nil, nil
	}
	jsonObj := new(Json)
	err := jsonObj.UnmarshalGraphQL(input)
	return jsonObj, err
}
