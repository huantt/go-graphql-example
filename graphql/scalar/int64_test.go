package graphql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64(t *testing.T) {
	num := new(Int64)
	err := num.UnmarshalGraphQL("123")
	assert.Error(t, err)
	err = num.UnmarshalGraphQL(123)
	assert.NoError(t, err)
}
