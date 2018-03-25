package scenario

import "testing"

type testingMock struct {
	errorOccurred bool
}

//noinspection SpellCheckingInspection
func (t *testingMock) Errorf(string, ...interface{})  {
	t.errorOccurred = true
}

func TestAssertPanics_ok(t *testing.T) {
	//Given:
	tc := "TestAssertPanics_ok"
	testingMock := &testingMock{}
	panickingFunction := func() { panic("panickingFunction") }

	//When:
	AssertPanics(testingMock, tc, panickingFunction)

	//Then:
	if testingMock.errorOccurred {
		t.Errorf("%v: panic was expected!", tc)
	}
}

func TestAssertPanics_nok(t *testing.T) {
	//Given:
	tc := "TestAssertPanics_nok"
	testingMock := &testingMock{}
	panickingFunction := func() {}

	//When:
	AssertPanics(testingMock, tc, panickingFunction)

	//Then:
	if !testingMock.errorOccurred {
		t.Errorf("%v: panic was not expected!", tc)
	}
}
