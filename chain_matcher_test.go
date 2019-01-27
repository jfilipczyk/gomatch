package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainMatcher(t *testing.T) {
	m := NewChainMatcher(
		[]ValueMatcher{
			NewNumberMatcher("@number@"),
			NewStringMatcher("@string@"),
		},
	)

	assert.True(t, m.CanMatch("@number@"), "expected to support @number@ pattern")
	assert.True(t, m.CanMatch("@string@"), "expected to support @string@ pattern")
	assert.False(t, m.CanMatch("@bool@"), "not expected to support @bool@ pattern")

	ok, err := m.Match("@number@", 123.)
	assert.True(t, ok, "expected to match number")
	assert.Nil(t, err, "expected to match number without any error")

	ok, err = m.Match("@string@", "some string")
	assert.True(t, ok, "expected to match string")
	assert.Nil(t, err, "expected to match string without any error")

	ok, err = m.Match("@bool@", true)
	assert.False(t, ok, "not expected to match bool")
	assert.EqualError(t, err, "none of matchers could be used")
}
