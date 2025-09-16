package contracts

type UserProvider[T any] interface {
	Get(id uint) (*T, error)
}
