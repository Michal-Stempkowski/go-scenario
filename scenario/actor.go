package scenario

import "testing"

type Actor interface {
	Enters()
	Leaves()
	Name() string
	T() *testing.T
	TestCase() string
	Enlist(cast Cast) Actor
}

type customActor struct {
	t              *testing.T
	tc             string
	enters, leaves func(Actor)
	name           string
}

func (c *customActor) Enters() {
	c.enters(c)
}

func (c *customActor) Leaves() {
	c.leaves(c)
}

func (c *customActor) Name() string {
	return c.name
}

func (c *customActor) T() *testing.T {
	return c.t
}

func (c *customActor) TestCase() string {
	return c.tc
}

func (c *customActor) Enlist(cast Cast) Actor {
	cast.Enlist(c)
	return c
}

func NewCustomActor(
	t *testing.T, tc string, enters, leaves func(Actor), name string) Actor {
	return &customActor{
		t:      t,
		tc:     tc,
		enters: enters,
		leaves: leaves,
		name:   name,
	}
}
