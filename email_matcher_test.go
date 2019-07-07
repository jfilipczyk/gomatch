package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var emailMatcherTests = []struct {
	desc   string
	v      interface{}
	ok     bool
	errMsg string
}{
	{
		"Should match email",
		"joe.doe@gmail.com",
		true,
		"",
	},
	{
		"Should match email with IP",
		"joe.doe@192.168.1.5",
		true,
		"",
	},
	{
		"Should match email with hostname without dot",
		"joe.doe@somehostname",
		true,
		"",
	},
	{
		"Should not match email with underscore",
		"joe.doe@my_mail.com",
		false,
		"expected email",
	},
	{
		"Should not match without hostname",
		"joe.doe@",
		false,
		"expected email",
	},
	{
		"Should not match without @",
		"joe.doe[at]gmail.com",
		false,
		"expected email",
	},
	{
		"Should not match user/box name",
		"@gmail.com",
		false,
		"expected email",
	},
	{
		"Should not match number",
		1234,
		false,
		"expected email",
	},
	{
		"Should not match slice",
		[]interface{}{"a", "b"},
		false,
		"expected email",
	},
}

func TestEmailMatcher(t *testing.T) {
	pattern := "@pattern@"

	for _, tt := range emailMatcherTests {
		m := NewEmailMatcher(pattern)
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
