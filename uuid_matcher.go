package gomatch

import (
	"errors"

	"github.com/google/uuid"
)

var errNotUUID = errors.New("expected UUID")

// A UUIDMatcher matches booleans.
type UUIDMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *UUIDMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching against given pattern.
func (m *UUIDMatcher) Match(p, v interface{}) (bool, error) {
	s, ok := v.(string)
	if !ok {
		return false, errNotUUID
	}
	_, err := uuid.Parse(s)
	if err != nil {
		return false, errNotUUID
	}
	return true, nil
}

// NewUUIDMatcher creates UUIDMatcher.
func NewUUIDMatcher(pattern string) *UUIDMatcher {
	return &UUIDMatcher{pattern}
}
