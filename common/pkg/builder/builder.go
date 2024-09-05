package builder

type Builder[T any] interface {
	MustBuild() T
	Build() (T, error)
}
