package scenario

import "testing"

func TestAssertPanics(t *testing.T) {
	tc := "TestAssertPanics"
	panickingFunction := func() { panic("panickingFunction") }
	AssertPanics(t, tc, panickingFunction)
}
