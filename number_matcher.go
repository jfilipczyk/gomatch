package gomatch

import "errors"

var errNotNumber = errors.New("expected number")

// A NumberMatcher matches float64.
// It expects float64 because json.Unmarshal uses float64 by default for numbers.
type NumberMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *NumberMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching against given pattern.
func (m *NumberMatcher) Match(p, v interface{}) (bool, error) {
	_, ok := v.(float64)
	if ok {
		return ok, nil
	}
	return ok, errNotNumber
}

// NewNumberMatcher creates NumberMatcher.
func NewNumberMatcher(pattern string) *NumberMatcher {
	return &NumberMatcher{pattern}
}
