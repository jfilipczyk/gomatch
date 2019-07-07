package gomatch

import (
	"errors"
	"regexp"
)

var emailRe = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var errNotEmail = errors.New("expected email")

// An EmailMatcher matches email
type EmailMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *EmailMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match performs value matching against given pattern.
func (m *EmailMatcher) Match(p, v interface{}) (bool, error) {
	s, ok := v.(string)
	if !ok {
		return false, errNotEmail
	}
	ok = emailRe.MatchString(s)
	if !ok {
		return false, errNotEmail
	}
	return true, nil
}

// NewEmailMatcher creates EmailMatcher.
func NewEmailMatcher(pattern string) *EmailMatcher {
	return &EmailMatcher{pattern}
}
