package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numberMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		// json package uses float64 when unmarshals to interface{}
		"Should match float64",
		100.,
		true,
		"",
	},
	{
		"Should not match string",
		"100",
		false,
		"expected number",
	},
	{
		"Should not match bool",
		true,
		false,
		"expected number",
	},
}

func TestNumberMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range numberMatcherTests {
		m := NewNumberMatcher(pattern)
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
