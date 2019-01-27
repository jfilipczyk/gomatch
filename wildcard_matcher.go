package gomatch

// A WildcardMatcher matches any value.
type WildcardMatcher struct {
	pattern string
}

// CanMatch returns true if pattern p can be handled
func (m *WildcardMatcher) CanMatch(p interface{}) bool {
	return isPattern(p, m.pattern)
}

// Match return true for any value
func (m *WildcardMatcher) Match(p, v interface{}) (bool, error) {
	return true, nil
}

// NewWildcardMatcher creates WildcardMatcher.
func NewWildcardMatcher(pattern string) *WildcardMatcher {
	return &WildcardMatcher{pattern}
}
