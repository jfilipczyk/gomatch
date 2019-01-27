package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var boolMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match true",
		true,
		true,
		"",
	},
	{
		"Should match false",
		false,
		true,
		"",
	},
	{
		"Should not match string",
		"false",
		false,
		"expected bool",
	},
}

func TestBoolMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range boolMatcherTests {
		m := NewBoolMatcher(pattern)
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
