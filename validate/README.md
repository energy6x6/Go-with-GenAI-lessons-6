# Task 2 — Email/phone validation (table-driven tests)

Implement `ValidateEmail` and/or `ValidatePhone` in `validate.go`
(check with your mentor whether one or both are required).

**Requirements**

- Extend `emailCases` and/or `phoneCases` in `validate_test.go` to at
  least **8 cases each**, including edge cases (empty string, missing
  separators, whitespace, unicode, boundary lengths, etc.).
- All cases must pass.

Run locally with:

```bash
go test -v ./validate/...
```

Each case is its own subtest (`t.Run`), so a single wrong case never
hides the others in the output.
