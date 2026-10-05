package validate

import "testing"

const minCases = 8

// The contract intentionally checks basic syntax, not RFC email deliverability.
var emailCases = []struct {
	name, input string
	want        bool
}{
	{"valid simple", "student@softserve.academy", true},
	{"missing at sign", "student-softserve.academy", false},
	{"empty string", "", false},
	{"missing local part", "@example.com", false},
	{"missing domain", "student@", false},
	{"multiple at signs", "a@b@c", false},
	{"space", "first last@example.com", false},
	{"leading space", " a@example.com", false},
	{"newline", "a@example.com\n", false},
	{"Unicode space", "a\u00a0@example.com", false},
	{"plus addressing", "a+tag@example.com", true},
	{"Unicode letters", "ім’я@пошта.укр", true},
	{"local domain", "student@localhost", true},
}

func TestValidateEmail(t *testing.T) {
	if len(emailCases) < minCases {
		t.Fatalf("need at least %d cases", minCases)
	}
	for _, tc := range emailCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidateEmail(tc.input); got != tc.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

var phoneCases = []struct {
	name, input string
	want        bool
}{
	{"valid with plus", "+380501234567", true},
	{"contains letters", "050-abc-4567", false},
	{"empty string", "", false},
	{"ten digits", "0501234567", true},
	{"fifteen digits", "+123456789012345", true},
	{"nine digits", "123456789", false},
	{"sixteen digits", "+1234567890123456", false},
	{"plus only", "+", false},
	{"double plus", "++380501234567", false},
	{"embedded plus", "380+501234567", false},
	{"spaces", "+380 501234567", false},
	{"separators", "050-123-4567", false},
	{"Unicode digits", "０５０１２３４５６７", false},
}

func TestValidatePhone(t *testing.T) {
	if len(phoneCases) < minCases {
		t.Fatalf("need at least %d cases", minCases)
	}
	for _, tc := range phoneCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidatePhone(tc.input); got != tc.want {
				t.Errorf("ValidatePhone(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
