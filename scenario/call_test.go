package scenario

import (
	"fmt"
	"testing"
)

func TestNewCall(t *testing.T) {
	//Given:
	tc := "TestNewCall"
	someString := "some string"
	someOtherArg := 5

	//When:
	uut := NewCall(someOtherArg, someString)

	//Then:
	a := uut.args
	if len(a) != 2 || a[0] != someOtherArg || a[1] != someString {
		t.Errorf("%v: unexpected arguments stored: %v", tc, a)
	}
}

func TestCall_Returns(t *testing.T) {
	//Given:
	tc := "TestCall_Returns"
	someArg := "some arg"
	uut := NewCall(someArg)
	if uut.returnValues != nil {
		t.Errorf(
			"%v: there should be no value; there is a: %v",
			tc, uut.returnValues,
		)
	}

	//When:
	uut.Returns(5, nil)

	//Then:
	r := uut.returnValues
	if len(r) != 2 || r[0] != 5 || r[1] != nil {
		t.Errorf("%v: unexpected return values stored: %v", tc, r)
	}
}

func TestCall_String(t *testing.T) {
	//Given:
	tc := "TestCall_String"
	someArg := "some arg"
	uut := NewCall(someArg)
	uut.Returns(5, nil)

	//When:
	s := fmt.Sprint(uut)

	//Then:
	if s != "Call[(some arg)->(5, <nil>)]" {
		t.Errorf("%v: unexpected string generated: %v", tc, s)
	}
}

func TestCall_Equal(t *testing.T) {
	equalityCheck := func(t2 *testing.T, a, b *Call, result bool) {
		if a.Equal(b) != result {
			message := "should"
			if !result {
				message = message + " not"
			}
			t2.Errorf("call %v %v be equal to %v", a, message, b)
		}
	}

	t.Run("TwoEmpty", func(t2 *testing.T) {
		equalityCheck(t2, NewCall(), NewCall(), true)
	})
	t.Run("ArgLenDiffer", func(t2 *testing.T) {
		equalityCheck(t2, NewCall(5), NewCall(), false)
	})
	t.Run("ArgDiffer", func(t2 *testing.T) {
		equalityCheck(t2, NewCall(5), NewCall(4), false)
	})
	t.Run("ReturnLenDiffer", func(t2 *testing.T) {
		a := NewCall(5)
		a.Returns(5)
		b := NewCall(5)
		b.Returns(5, nil)
		equalityCheck(t2, a, b, false)
	})
	t.Run("ReturnValuesDiffer", func(t2 *testing.T) {
		a := NewCall(5)
		a.Returns(5)
		b := NewCall(5)
		b.Returns(4)
		equalityCheck(t2, a, b, false)
	})
	t.Run("TwoSame", func(t2 *testing.T) {
		a := NewCall(5)
		a.Returns(9)
		b := NewCall(5)
		b.Returns(9)
		equalityCheck(t2, a, b, true)
	})
	t.Run("ReturnsNothingSyntacticSugar", func(t2 *testing.T) {
		a := NewCall(5)
		a.Returns()
		b := NewCall(5)
		b.ReturnsNothing()
		equalityCheck(t2, a, b, true)
	})
}

func TestCall_Describe(t *testing.T) {
	//Given:
	tc := "TestCall_String"
	someArg := "some arg"
	uut := NewCall(someArg)
	uut.Returns(5, nil).Describe("Description")

	//When:
	s := fmt.Sprint(uut)

	//Then:
	if s != "Call[(some arg)->(5, <nil>) //Description]" {
		t.Errorf("%v: unexpected string generated: %v", tc, s)
	}
}
