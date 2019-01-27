package gomatch

import "errors"

var errMatcherNotFound = errors.New("none of matchers could be used")

// A ChainMatcher allows to chain multiple value matchers
type ChainMatcher struct {
	matchers []ValueMatcher
}

// CanMatch returns true if pattern p can be handled by any of internal matchers
func (m *ChainMatcher) CanMatch(p interface{}) bool {
	for _, m := range m.matchers {
		if m.CanMatch(p) {
			return true
		}
	}
	return false
}

// Match performs value matching against given pattern.
// It iterates through internal matchers and uses first which can handle given pattern.
func (m *ChainMatcher) Match(p, v interface{}) (bool, error) {
	for _, m := range m.matchers {
		if !m.CanMatch(p) {
			continue
		}
		return m.Match(p, v)

	}
	return false, errMatcherNotFound
}

// NewChainMatcher creates ChainMatcher.
func NewChainMatcher(matchers []ValueMatcher) *ChainMatcher {
	return &ChainMatcher{matchers}
}
