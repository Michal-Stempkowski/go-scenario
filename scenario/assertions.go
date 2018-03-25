package scenario

//noinspection SpellCheckingInspection
type TestingT interface {
	Errorf(string, ...interface{})
}

func AssertPanics(t TestingT, tc string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v: Panic expected, but none happened!", tc)
		}
	}()
	f()
}
