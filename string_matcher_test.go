package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var stringMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match string",
		"some valid string",
		true,
		"",
	},
	{
		"Should not match number",
		1234,
		false,
		"expected string",
	},
	{
		"Should not match slice",
		[]interface{}{"a", "b"},
		false,
		"expected string",
	},
}

func TestStringMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range stringMatcherTests {
		m := NewStringMatcher(pattern)
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
