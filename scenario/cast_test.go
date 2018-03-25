package scenario

import (
	"testing"
)

type actorMock struct {
	entersCount, leavesCount int
	name                     string
	t                        *testing.T
	tc                       string
}

func (a *actorMock) Enters() {
	a.entersCount++
}

func (a *actorMock) Leaves() {
	a.leavesCount++
}

func (a *actorMock) Name() string {
	return a.name
}

func (a *actorMock) T() *testing.T {
	return a.t
}

func (a *actorMock) TestCase() string {
	return a.tc
}

func (a *actorMock) Enlist(cast Cast) Actor {
	return a
}

func newActorMock(t *testing.T, tc string, name string) *actorMock {
	return &actorMock{name: name, t: t, tc: tc}
}

func TestCast_Enlist(t *testing.T) {
	//Given:
	tc := "TestCast_Enlist"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	if len(uut.actors) != 0 {
		t.Errorf("%v: There should be no cast!", tc)
	}

	//When:
	uut.Enlist(actor)

	//Then:
	if len(uut.actors) != 1 || uut.actors[0] != actor {
		t.Errorf("%v: Actor should be enlisted!", tc)
	}
}

func TestCast_StartPlay(t *testing.T) {
	//Given:
	tc := "TestCast_StartPlay"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)

	//When:
	uut.StartPlay()

	//Then:
	if actor.entersCount != 1 {
		t.Errorf("%v: Actor should enter exactly once!", tc)
	}
}

func TestCast_StartPlayCalledTwice(t *testing.T) {
	//Given:
	tc := "TestCast_StartPlay"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)
	uut.StartPlay()

	//When/Then:
	AssertPanics(t, tc, uut.StartPlay)
	if actor.entersCount != 1 {
		t.Errorf("%v: Actor should enter exactly once!", tc)
	}
}

func TestCast_EndPlayWhenNoPlayInProgress(t *testing.T) {
	//Given:
	tc := "TestCast_EndPlayWhenNoPlayInProgress"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)

	//When/Then:
	AssertPanics(t, tc, uut.EndPlay)
	if actor.leavesCount != 0 {
		t.Errorf("%v: Actor should not leave if not entered!", tc)
	}
}

func TestCast_EndPlayCanBeCalledAfterStartPlay(t *testing.T) {
	//Given:
	tc := "TestCast_EndPlayCanBeCalledAfterStartPlay"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)
	uut.StartPlay()

	//When:
	uut.EndPlay()

	//Then:
	if actor.entersCount != 1 {
		t.Errorf("%v: Actor should enter exactly once!", tc)
	}
	if actor.leavesCount != 1 {
		t.Errorf("%v: Actor should leave exactly once!", tc)
	}
}

func TestCast_StartPlayCanBeCalledAfterEndPlay(t *testing.T) {
	//Given:
	tc := "TestCast_StartPlayCanBeCalledAfterEndPlay"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)
	uut.StartPlay()
	uut.EndPlay()

	//When:
	uut.StartPlay()

	//Then:
	if actor.entersCount != 2 {
		t.Errorf("%v: Actor should enter exactly once!", tc)
	}
	if actor.leavesCount != 1 {
		t.Errorf("%v: Actor should leave exactly once!", tc)
	}
}

func TestCast_Play(t *testing.T) {
	//Given:
	tc := "TestCast_Play"
	uut := NewCast()
	actor := newActorMock(t, tc, "mocked actor")
	uut.Enlist(actor)

	func() {
		defer uut.Play()()

		//Then:
		if actor.entersCount != 1 {
			t.Errorf("%v: Actor should enter now!", tc)
		}
		if actor.leavesCount != 0 {
			t.Errorf("%v: Actor's leave should be deferred!", tc)
		}
	}()

	//Then:
	if actor.entersCount != 1 {
		t.Errorf("%v: Actor should enter exactly once!", tc)
	}
	if actor.leavesCount != 1 {
		t.Errorf("%v: Actor should leave exactly once!", tc)
	}
}
