package config

import "testing"

// These cases were generated from the signature, Config fields and validation
// rules before implementing ValidateConfig. There is no human-written baseline
// to compare with in this AI-assisted solution. The boundary pairs 1023/1024 and
// 65535/65536 plus whitespace and Unicode inputs make the exact contract explicit.
// No incorrect expectations were found when reviewing these cases against it.
func TestValidateConfig(t *testing.T) {
	cases := []struct {
		name    string
		port    int
		env     string
		wantErr bool
	}{
		{"development lower bound", 1024, "development", false},
		{"staging", 8080, "staging", false},
		{"production upper bound", 65535, "production", false},
		{"below lower bound", 1023, "development", true},
		{"above upper bound", 65536, "development", true},
		{"zero", 0, "development", true},
		{"privileged port", 80, "development", true},
		{"negative port", -1, "development", true},
		{"large port", 99999, "development", true},
		{"empty environment", 8080, "", true},
		{"unknown environment", 8080, "test", true},
		{"uppercase", 8080, "Production", true},
		{"whitespace", 8080, "production ", true},
		{"newline", 8080, "production\n", true},
		{"Unicode lookalike", 8080, "prоduction", true},
		{"zero value", 0, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateConfig(Config{ServerPort: tc.port, Environment: tc.env})
			if (err != nil) != tc.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
