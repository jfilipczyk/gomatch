package gomatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonMatcherTests = []struct {
	desc   string
	p      string
	v      string
	ok     bool
	errMsg string
}{
	{
		"Should fail if invalid JSON pattern given",
		`{"foo":}`,
		`{"foo": "bar"}`,
		false,
		"invalid JSON pattern",
	},
	{
		"Should fail if invalid JSON given",
		`{"foo": "bar"}`,
		`{"foo":}`,
		false,
		"invalid JSON",
	},
	{
		"Should succeed if strings are equal",
		`"John"`,
		`"John"`,
		true,
		"",
	},
	{
		"Should succeed if numbers are equal",
		`123`,
		`123`,
		true,
		"",
	},
	{
		"Should succeed if bools are equal",
		`true`,
		`true`,
		true,
		"",
	},
	{
		"Should fail if types are not equal",
		`"John"`,
		`true`,
		false,
		"types are not equal",
	},
	{
		"Should fail if values are not equal",
		`100`,
		`200`,
		false,
		"values are not equal",
	},
	{
		"Should succeed if objects are equal",
		`
		{
			"id": 123,
			"name": "John Smith"
		}
		`,
		`
		{
			"id": 123,
			"name": "John Smith"
		}
		`,
		true,
		"",
	},
	{
		"Should succeed if objects are equal but with different key order",
		`
		{
			"id": 123,
			"name": "John Smith"
		}
		`,
		`
		{
			"name": "John Smith",
			"id": 123
		}
		`,
		true,
		"",
	},
	{
		"Should succeed if arrays are equal",
		"[1,2,3]",
		"[1,2,3]",
		true,
		"",
	},
	{
		"Should fail if array values in different order",
		"[1,2,3]",
		"[1,3,2]",
		false,
		"values are not equal at path: [1]",
	},
	{
		"Should fail if has same keys but values differ",
		`
		{
			"id": 123,
			"name": "John Smith"
		}
		`,
		`
		{
			"id": 999,
			"name": "John Smith"
		}
		`,
		false,
		"values are not equal at path: id",
	},
	{
		"Should succeed if nested objects are equal",
		`
		{
			"id": 123,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "work",
					"phone": "222-222-222"
				}
			]
		}
		`,
		`
		{
			"id": 123,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "work",
					"phone": "222-222-222"
				}
			]
		}
		`,
		true,
		"",
	},
	{
		"Should fail if nested objects are not equal",
		`
		{
			"id": 123,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "work",
					"phone": "222-222-222"
				}
			]
		}
		`,
		`
		{
			"id": 123,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "mobile",
					"phone": "222-222-222"
				}
			]
		}
		`,
		false,
		"values are not equal at path: phones[1].type",
	},
	{
		"Should succeed if values matches patterns",
		`
		{
			"id": "@wildcard@",
			"name": "@wildcard@",
			"phones": [
				{
					"type": "home",
					"phone": "@wildcard@"
				},
				"@wildcard@"
			]
		}
		`,
		`
		{
			"id": 123,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "office",
					"phone": "222-222-222"
				}
			]
		}
		`,
		true,
		"",
	},
	{
		"Should fail if object does not have all expected keys",
		`
		{
			"id": 1,
			"name": "John Smith"
		}
		`,
		`
		{
			"id": 1
		}
		`,
		false,
		`expected key "name"`,
	},
	{
		"Should fail if object has unexpected keys",
		`
		{
			"id": 1
		}
		`,
		`
		{
			"id": 1,
			"name": "John Smith"
		}
		`,
		false,
		"unexpected key",
	},
	{
		"Should succeed if object has unexpected keys but unbounded pattern was used",
		`
		{
			"id": 1,
			"@...@": ""
		}
		`,
		`
		{
			"id": 1,
			"name": "John Smith"
		}
		`,
		true,
		"",
	},
	{
		"Should fail if array has unexpected extra values",
		"[1,2,3]",
		"[1,2,3,4]",
		false,
		"arrays sizes are not equal",
	},
	{
		"Should fail if array misses some values",
		"[1,2,3]",
		"[1,2]",
		false,
		"arrays sizes are not equal",
	},
	{
		"Should fail if array misses some values but unbounded pattern was used",
		`[1,2,"@...@"]`,
		"[1]",
		false,
		"arrays sizes are not equal",
	},
	{
		"Should succeed if array has unexpected extra values but unbounded pattern was used",
		`[1,2,3,"@...@"]`,
		"[1,2,3,4]",
		true,
		"",
	},
	{
		"Should succeed if nested object has extra values but unbounded patterns were used",
		`
		{
			"name": "John Smith",
			"phones": [
				{
					"phone": "111-111-111",
					"@...@": ""
				},
				"@...@"
			],
			"@...@": ""
		}
		`,
		`
		{
			"id": 1,
			"name": "John Smith",
			"phones": [
				{
					"type": "home",
					"phone": "111-111-111"
				},
				{
					"type": "office",
					"phone": "222-222-222"
				}
			]
		}
		`,
		true,
		"",
	},
	{
		"Should fail if unknown value pattern was used (only @wildcard@ was setup in this suite)",
		`
		{
			"id": "@wildcard@",
			"name": "@string@"
		}
		`,
		`
		{
			"id": 1,
			"name": "John Smith"
		}
		`,
		false,
		"values are not equal at path: name",
	},
}

func TestJSONMatcher(t *testing.T) {
	for _, tt := range jsonMatcherTests {
		m := NewJSONMatcher(NewWildcardMatcher(patternWildcard))
		ok, err := m.Match(tt.p, tt.v)

		t.Logf(tt.desc)
		if tt.ok {
			assert.Nil(t, err)
			assert.True(t, ok)
		} else {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.errMsg)
			assert.False(t, ok)
		}
	}
}

func TestJSONMatcherWithDefaultMatchers(t *testing.T) {
	p := `
	{
		"id": "@number@",
		"uuid": "@uuid@",
		"name": "@string@",
		"isActive": "@bool@",
		"createdAt": "@wildcard@",
		"phones": "@array@",
		"email": "@email@",
		"@...@": ""
	}
	`
	v := `
	{
		"id": 1,
		"uuid": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"name": "John Smith",
		"isActive": true,
		"createdAt": "2018-05-27T12:00:00Z",
		"phones": [
			{
				"type": "home",
				"phone": "111-111-111"
			},
			{
				"type": "office",
				"phone": "222-222-222"
			}
		],
		"email": "john.smith@gmail.com",
		"isVip": false
	}
	`

	m := NewDefaultJSONMatcher()
	ok, err := m.Match(p, v)

	assert.Nil(t, err)
	assert.True(t, ok)
}
