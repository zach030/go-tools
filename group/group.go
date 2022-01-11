package group

type Group struct {
	actors []actor
}

func (g *Group) Add(exe func() error, interrupt func(error)) {
	g.actors = append(g.actors, actor{execute: exe, interrupt: interrupt})
}

func (g *Group) Run() error {
	if len(g.actors) == 0 {
		return nil
	}
	errors := make(chan error, len(g.actors))
	// run each actors
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}
	// wait for the first error
	err := <-errors
	// signal all actors to stop
	for _, a := range g.actors {
		a.interrupt(err)
	}
	// wait for all actors to stop
	return err
}

type actor struct {
	execute   func() error
	interrupt func(error)
}
