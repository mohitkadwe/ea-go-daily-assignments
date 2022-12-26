package quickcash

type Withdrawable interface {
	CanWithDraw(amount float64) (bool, error)
	WithDraw(amount float64) float64
	GetIdentifier() string
}
