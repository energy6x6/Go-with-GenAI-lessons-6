// Package validate provides simple syntactic validators for common
// user-input formats.
//
// Homework — Task 2 (Lesson 6: File I/O, JSON and Testing):
// Implement ValidateEmail and/or ValidatePhone below (your mentor may
// ask for just one) and extend the test tables in validate_test.go to
// at least 8 cases each, including edge cases.
package validate

// ValidateEmail reports whether s is a syntactically valid email address.
//
// TODO: implement this function. At minimum it should:
//   - reject the empty string,
//   - require exactly one "@" with a non-empty local part and domain part,
//   - reject values containing whitespace.
//
// Document any additional decisions you make (e.g. how you handle a
// trailing dot, consecutive dots, or unicode characters) in a comment
// here, and add matching test cases in validate_test.go.
func ValidateEmail(s string) bool {
	// TODO: implement me
	return false
}

// ValidatePhone reports whether s is a syntactically valid phone number.
//
// TODO: implement this function. Decide which format(s) you accept
// (e.g. "+380501234567", "050-123-4567") and document your decision
// here. At minimum it should reject the empty string and any value
// containing letters.
func ValidatePhone(s string) bool {
	// TODO: implement me
	return false
}
