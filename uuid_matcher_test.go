package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uuidMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match UUID",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		true,
		"",
	},
	{
		"Should not match invalid UUID",
		"6ba7b810-9dad-XXXX-80b4-00c04fd430c8",
		false,
		"expected UUID",
	},
}

func TestUUIDMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range uuidMatcherTests {
		m := NewUUIDMatcher(pattern)
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
