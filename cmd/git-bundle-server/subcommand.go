package main

type Subcommand interface {
	subcommand() string
	run(args []string) error
}

func all() []Subcommand {
	return []Subcommand{
		Delete{},
		Init{},
		Start{},
		Stop{},
		Update{},
	}
}
