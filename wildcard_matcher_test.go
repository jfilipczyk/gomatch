package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wildcardMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match everything - string",
		"some string",
		true,
		"",
	},
	{
		"Should match everything - array",
		[]interface{}{1, 2, 3},
		true,
		"",
	},
	{
		"Should match everything - number",
		100.,
		true,
		"",
	},
	{
		"Should match everything - null",
		nil,
		true,
		"",
	},
	{
		"Should match everything - map",
		map[string]interface{}{"key": "value"},
		true,
		"",
	},
}

func TestWildcardMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range wildcardMatcherTests {
		m := NewWildcardMatcher(pattern)
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
