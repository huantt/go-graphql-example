package schema_test

import (
	"github.com/huantt/go-graphql-sample/graphql/schema"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	s, err := schema.String()

	require.NoError(t, err)
	require.NotEmpty(t, s)
}
