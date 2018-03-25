package scenario

import "testing"

func AssertPanics(t *testing.T, tc string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v: Panic expected, but none happened!", tc)
		}
	}()
	f()
}
