package gomatch

import "errors"

var errNotArray = errors.New("expected array")

// An ArrayMatcher matches []interface{}.
type ArrayMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *ArrayMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching agains given pattern.
func (m *ArrayMatcher) Match(p, v interface{}) (bool, error) {
	_, ok := v.([]interface{})
	if ok {
		return ok, nil
	}
	return ok, errNotArray
}

// NewArrayMatcher creates ArrayMatcher.
func NewArrayMatcher(pattern string) *ArrayMatcher {
	return &ArrayMatcher{pattern}
}
