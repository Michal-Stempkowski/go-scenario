package scenario

import "testing"

type actorTestSuite struct {
	entersCallCount, leavesCallCount int
	tc                               string
	name                             string
	t                                *testing.T
	castMock                         *fakeCastForEnlist
}

func (a *actorTestSuite) enters(actor Actor) {
	a.entersCallCount++
}

func (a *actorTestSuite) leaves(actor Actor) {
	a.leavesCallCount++
}

func (a *actorTestSuite) expectEntersCallCount(count int) {
	a.expectCallCount("enters", a.entersCallCount, count)
}

func (a *actorTestSuite) expectLeavesCallCount(count int) {
	a.expectCallCount("leaves", a.leavesCallCount, count)
}

func (a *actorTestSuite) expectStringEqual(received, expected string) {
	if received != expected {
		a.t.Errorf(
			"%v: Expected name: %v, received: %v",
			a.tc, expected, received,
		)
	}
}

func (a *actorTestSuite) expectTEqual(received, expected *testing.T) {
	if received != expected {
		a.t.Errorf(
			"%v: Expected name: %v, received: %v",
			a.tc, expected, received,
		)
	}
}

func (a *actorTestSuite) expectCallCount(
	name string, received, expected int) {
	if received != expected {
		a.t.Errorf(
			"%v: Expected %v call count: %v, received: %v",
			a.tc, name, expected, received,
		)
	}
}

func newActorTemplateTestSuite(tc string) *actorTestSuite {
	return &actorTestSuite{
		tc:       tc,
		name:     "test actor",
		castMock: &fakeCastForEnlist{},
	}
}

type fakeCastForEnlist struct {
	enlistCalls int
}

func (f *fakeCastForEnlist) Enlist(actor Actor) {
	f.enlistCalls++
}

func (f *fakeCastForEnlist) StartPlay() {}

func (f *fakeCastForEnlist) EndPlay() {}

func (f *fakeCastForEnlist) Play() func() {
	return f.EndPlay
}

func TestNewCustomActorEnters(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorEnters")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//When:
	uut.Enters()

	//Then:
	ts.expectEntersCallCount(1)
}

func TestNewCustomActorLeaves(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorLeaves")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//When:
	uut.Leaves()

	//Then:
	ts.expectLeavesCallCount(1)
}

func TestNewCustomActorName(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorName")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//Then:
	ts.expectStringEqual(uut.Name(), ts.name)
}

func TestCustomActor_T(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorName")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//Then:
	ts.expectTEqual(uut.T(), t)
}

func TestCustomActor_TestCase(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorName")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//Then:
	ts.expectStringEqual(uut.TestCase(), ts.tc)
}

func TestCustomActor_Enlist(t *testing.T) {
	//Given:
	ts := newActorTemplateTestSuite("TestNewCustomActorName")
	uut := NewCustomActor(
		t, ts.tc, ts.enters, ts.leaves, ts.name,
	)

	//When:
	uut.Enlist(ts.castMock)

	//Then:
	ts.expectCallCount(ts.tc, ts.castMock.enlistCalls, 1)
}
