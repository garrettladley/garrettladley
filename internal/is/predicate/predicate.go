package predicate

type Predicate byte

const (
	Even Predicate = iota
	Odd
	// TODO: prime? or more...?
)
