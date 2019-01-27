package gomatch

import "errors"

var errNotString = errors.New("expected string")

// A StringMatcher matches any string
type StringMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled.
func (m *StringMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching agains given pattern.
func (m *StringMatcher) Match(p, v interface{}) (bool, error) {
	_, ok := v.(string)
	if ok {
		return ok, nil
	}
	return ok, errNotString
}

// NewStringMatcher creates StringMatcher.
func NewStringMatcher(pattern string) *StringMatcher {
	return &StringMatcher{pattern}
}
