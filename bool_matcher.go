package gomatch

import "errors"

var errNotBool = errors.New("expected bool")

// A BoolMatcher matches booleans.
type BoolMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *BoolMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching against given pattern.
func (m *BoolMatcher) Match(p, v interface{}) (bool, error) {
	_, ok := v.(bool)
	if ok {
		return ok, nil
	}
	return ok, errNotBool
}

// NewBoolMatcher creates BoolMatcher.
func NewBoolMatcher(pattern string) *BoolMatcher {
	return &BoolMatcher{pattern}
}
