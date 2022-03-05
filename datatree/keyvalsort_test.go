package datatree

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValSort(t *testing.T) {
	kvs := keyValList{
		{
			key: "abc",
		},
		{
			key: "123",
		},
		{
			key: "def",
		},
	}

	sort.Sort(kvs)

	assert.Equal(t, "123", kvs[0].key)
	assert.Equal(t, "abc", kvs[1].key)
	assert.Equal(t, "def", kvs[2].key)
}
