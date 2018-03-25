package scenario

type Cast interface {
	Enlist(actor Actor)
	StartPlay()
	EndPlay()
	Play() func()
}

type CastType struct {
	actors         []Actor
	playInProgress bool
}

func NewCast(actors ...Actor) *CastType {
	return &CastType{actors: actors}
}

func (c *CastType) Enlist(a Actor) {
	c.actors = append(c.actors, a)
}

func (c *CastType) StartPlay() {
	if c.playInProgress {
		panic("Play is already in progress!")
	}
	c.playInProgress = true
	for _, a := range c.actors {
		a.Enters()
	}
}

func (c *CastType) EndPlay() {
	if !c.playInProgress {
		panic("Play has not started yet!")
	}
	c.playInProgress = false

	for _, a := range c.actors {
		a.Leaves()
	}
}

func (c *CastType) Play() func() {
	c.StartPlay()
	return c.EndPlay
}
