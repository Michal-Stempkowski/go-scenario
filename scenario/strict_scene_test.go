package scenario

import "testing"

func TestStrictScene_Call(t *testing.T) {
	//Given:
	tc := "TestStrictScene_Call"
	someArg := 5
	uut := NewStrictScene()
	if uut.expectedCalls != nil {
		t.Errorf(
			"%v: No c should be expected. Received: %v",
			tc, uut.expectedCalls,
		)
	}

	//When:
	uut.Call(someArg)

	//Then:
	c := uut.expectedCalls
	if len(c) != 1 || len(c[0].args) != 1 || !c[0].Equal(NewCall(someArg)) {
		t.Errorf("%v: expected calls differ: %v", tc, c)
	}
}

func TestStrictSceneNoFailures(t *testing.T) {
	tc := "TestStrictSceneNoFailures"
	uut := NewStrictScene()
	f := uut.Summarize()
	if len(f) != 0 {
		t.Errorf("%v: no failures expected but received: %v", tc, f)
	}
}

//func TestStrictSceneCallOmitted(t *testing.T) {
//	tc := "TestStrictSceneCallOmitted"
//	uut := NewStrictScene()
//	uut.Call(5).Returns(4)
//	f := uut.Summarize()
//	if len(f) != 0 {
//		t.Errorf("%v: no failures expected but received: %v", tc, f)
//	}
//}
