package graphql

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
)

type Int64 int64

func (*Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

func (m *Int64) UnmarshalGraphQL(input interface{}) error {
	if input == nil {
		return nil
	}
	switch inputValue := input.(type) {
	case uint64, uint32, uint, int64, int32, int:
		*m = Int64(reflect.ValueOf(inputValue).Convert(reflect.TypeOf(*m)).Int())
	case float64:
		*m = Int64(int64(inputValue))
	default:
		return fmt.Errorf("input type is %T - It must be int64/int32/int/float64", input)
	}
	return nil
}

func NewInt64[T constraints.Float | constraints.Integer](input T) (*Int64, error) {
	num := new(Int64)
	err := num.UnmarshalGraphQL(input)
	return num, err
}

func (m *Int64) ToInt64() int64 {
	return int64(*m)
}
