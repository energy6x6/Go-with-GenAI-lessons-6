// Package validate provides basic syntax checks, not deliverability checks.
package validate

import (
	"strings"
	"unicode"
)

// ValidateEmail requires one @, nonempty parts and no Unicode whitespace.
// Unicode letters and local domains are accepted; full RFC validation is out of scope.
func ValidateEmail(s string) bool {
	if strings.Count(s, "@") != 1 || strings.ContainsFunc(s, unicode.IsSpace) {
		return false
	}
	local, domain, _ := strings.Cut(s, "@")
	return local != "" && domain != ""
}

// ValidatePhone accepts 10–15 ASCII digits with an optional leading +.
// Spaces, separators, letters and Unicode digits are rejected.
func ValidatePhone(s string) bool {
	s = strings.TrimPrefix(s, "+")
	if len(s) < 10 || len(s) > 15 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
