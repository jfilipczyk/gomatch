// Package gomatch provides types for pattern based JSON matching.
//
// It provides JSONMatcher type which performs deep comparison of two JSON strings.
// JSONMatcher may be created with a set of ValueMatcher implementations.
// A ValueMatcher is used to make comparison less strict than a regular value comparison.
//
// Use NewDefaultJSONMatcher to create JSONMatcher with a chain of all available ValueMatcher implementations.
//
// Basic usage:
//
//  actual := `
//  {
//  	"id": 351,
//  	"name": "John Smith",
//  	"address": {
//  		"city": "Boston"
//  	}
//  }
//  `
//  expected := `
//  {
//  	"id": "@number@",
//  	"name": "John Smith",
//  	"address": {
//  		"city": "@string@"
//  	}
//  }
//  `
//
//  m := gomatch.NewDefaultJSONMatcher()
//  ok, err := m.Match(expected, actual)
//  if ok {
//  	fmt.Printf("actual JSON matches expected JSON")
//  } else {
//  	fmt.Printf("actual JSON does not match expected JSON: %s", err.Error())
//  }
//
// Use NewJSONMatcher to create JSONMatcher with a custom ValueMatcher implementation.
// Use ChainMatcher to chain multiple ValueMacher implementations.
//
//  m := gomatch.NewJSONMatcher(
//  	NewChainMatcher(
//  		[]ValueMatcher{
//  			NewStringMatcher("@string@"),
//  			NewNumberMatcher("@number@"),
//  		},
//  	)
//  );
//
package gomatch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	errInvalidJSON        = errors.New("invalid JSON")
	errInvalidJSONPattern = errors.New("invalid JSON pattern")
	errTypesNotEqual      = errors.New("types are not equal")
	errValuesNotEqual     = errors.New("values are not equal")
	errArraysLenNotEqual  = errors.New("arrays sizes are not equal")
	errUnexpectedKey      = errors.New("unexpected key")
)

const (
	patternString    = "@string@"
	patternNumber    = "@number@"
	patternBool      = "@bool@"
	patternArray     = "@array@"
	patternUUID      = "@uuid@"
	patternEmail     = "@email@"
	patternWildcard  = "@wildcard@"
	patternUnbounded = "@...@"
)

// A ValueMatcher interface should be implemented by any matcher used by JSONMatcher.
type ValueMatcher interface {
	// CanMatch returns true if given pattern can be handled by value matcher implementation.
	CanMatch(p interface{}) bool

	// Match performs the matching of given value v.
	// It also expects pattern p so implementation may handle multiple patterns or some DSL.
	Match(p, v interface{}) (bool, error)
}

// NewDefaultJSONMatcher creates JSONMatcher with default chain of value matchers.
// Default chain contains:
//
// - StringMatcher handling "@string@" pattern
//
// - NumberMatcher handling "@number@" pattern
//
// - BoolMatcher handling "@bool@" pattern
//
// - ArrayMatcher handling "@array@" pattern
//
// - UUIDMatcher handling "@uuid@" pattern
//
// - EmailMatcher handling "@email@" pattern
//
// - WildcardMatcher handling "@wildcard@" pattern
//
func NewDefaultJSONMatcher() *JSONMatcher {
	return NewJSONMatcher(
		NewChainMatcher(
			[]ValueMatcher{
				NewStringMatcher(patternString),
				NewNumberMatcher(patternNumber),
				NewBoolMatcher(patternBool),
				NewArrayMatcher(patternArray),
				NewUUIDMatcher(patternUUID),
				NewEmailMatcher(patternEmail),
				NewWildcardMatcher(patternWildcard),
			},
		))
}

// NewJSONMatcher creates JSONMatcher with given value matcher.
func NewJSONMatcher(matcher ValueMatcher) *JSONMatcher {
	return &JSONMatcher{matcher}
}

// A JSONMatcher provides Match method to match two JSONs with pattern matching support.
type JSONMatcher struct {
	valueMatcher ValueMatcher
}

// Match performs deep match of given JSON with an expected JSON pattern.
//
// It traverses expected JSON pattern and checks if actual JSON has expected values.
// When traversing it checks if expected value is a pattern supported by internal ValueMatcher.
// In such case it uses the ValueMatcher to match actual value otherwise it compares expected
// value with actual value.
//
// Expected JSON pattern example:
//  {
//  	"id": "@number@",
//  	"name": "John Smith",
//  	"address": {
//  		"city": "@string@"
//  	}
//  }
//
// Matching actual JSON:
//  {
//  	"id": 351,
//  	"name": "John Smith",
//  	"address": {
//  		"city": "Boston"
//  	}
//  }
//
// In above example we assume that ValueMatcher supports "@number@" and "@string@" patterns,
// otherwise matching will fail.
//
// Besides value patterns JSONMatcher supports an "unbounded pattern" - "@...@".
// It can be used at the end of an array to allow any extra array elements:
//
//  [
//  	"John Smith",
//  	"Joe Doe",
//  	"@...@"
//  ]
//
// It can be used at the end of an object to allow any extra keys:
//
//  {
//  	"id": 351,
//  	"name": "John Smith",
//  	"@...@": ""
//  }
//
// When matching fails then error message contains a path to invalid value.
func (m *JSONMatcher) Match(expectedJSON, actualJSON string) (bool, error) {
	var expected, actual interface{}
	err := json.Unmarshal([]byte(expectedJSON), &expected)
	if err != nil {
		return false, errInvalidJSONPattern
	}
	err = json.Unmarshal([]byte(actualJSON), &actual)
	if err != nil {
		return false, errInvalidJSON
	}
	path, err := m.deepMatch(expected, actual)
	if err != nil {
		if len(path) > 0 {
			err = fmt.Errorf("%s at path: %s", err.Error(), pathToString(path))
		}
		return false, err
	}
	return true, nil
}

func (m *JSONMatcher) deepMatch(expected interface{}, actual interface{}) ([]interface{}, error) {
	var path []interface{}
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) && !m.valueMatcher.CanMatch(expected) {
		return path, errTypesNotEqual
	}

	switch expected.(type) {
	case []interface{}:
		return m.deepMatchArray(expected.([]interface{}), actual.([]interface{}))

	case map[string]interface{}:
		return m.deepMatchMap(expected.(map[string]interface{}), actual.(map[string]interface{}))

	default:
		return m.matchValue(expected, actual)
	}
}

func (m *JSONMatcher) deepMatchArray(expected, actual []interface{}) ([]interface{}, error) {
	var path []interface{}
	unbounded := false
	for i, v := range expected {
		if isUnbounded(v) {
			unbounded = true
			break
		}
		if i == len(actual) {
			 break
		}
		keyPath, err := m.deepMatch(v, actual[i])
		if err != nil {
			return append(keyPath, i), err
		}
	}
	if !unbounded && len(expected) != len(actual) {
		return path, errArraysLenNotEqual
	}
	return path, nil
}

func (m *JSONMatcher) deepMatchMap(expected, actual map[string]interface{}) ([]interface{}, error) {
	var path []interface{}
	unbounded := false
	for k, v1 := range expected {
		if isUnbounded(k) {
			unbounded = true
			continue
		}
		v2, ok := actual[k]
		if !ok {
			return path, fmt.Errorf(`expected key "%s"`, k)
		}
		keyPath, err := m.deepMatch(v1, v2)
		if err != nil {
			return append(keyPath, k), err
		}
	}
	if !unbounded && len(expected) != len(actual) {
		return path, errUnexpectedKey
	}
	return path, nil
}

func (m *JSONMatcher) matchValue(expected, actual interface{}) ([]interface{}, error) {
	var path []interface{}
	if m.valueMatcher.CanMatch(expected) {
		_, err := m.valueMatcher.Match(expected, actual)
		return path, err
	}
	if expected != actual {
		return path, errValuesNotEqual
	}
	return path, nil
}

func pathToString(path []interface{}) string {
	var b bytes.Buffer
	for i := len(path) - 1; i > -1; i-- {
		v := path[i]
		switch v.(type) {
		case int:
			b.WriteString(fmt.Sprintf("[%d]", v.(int)))
			break
		default:
			if b.Len() > 0 {
				b.WriteRune('.')
			}
			b.WriteString(v.(string))
			break
		}
	}
	return b.String()
}

func isUnbounded(p interface{}) bool {
	return isPattern(p, patternUnbounded)
}

func isPattern(p interface{}, pattern string) bool {
	ps, ok := p.(string)
	return ok && ps == pattern
}
