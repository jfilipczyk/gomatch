package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arrayMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match slice",
		[]interface{}{1, 2, 3},
		true,
		"",
	},
	{
		"Should match empty slice",
		[]interface{}{},
		true,
		"",
	},
	{
		"Should not match string",
		"some string",
		false,
		"expected array",
	},
	{
		"Should not match nil",
		nil,
		false,
		"expected array",
	},
}

func TestArrayMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range arrayMatcherTests {
		m := NewArrayMatcher(pattern)
		assert.True(t, m.CanMatch(pattern), "expected to support pattern")

		t.Logf(tt.desc)

		ok, err := m.Match(pattern, tt.v)

		if tt.ok {
			assert.True(t, ok)
			assert.Nil(t, err)
		} else {
			assert.False(t, ok)
			assert.EqualError(t, err, tt.errMsg)
		}
	}
}
