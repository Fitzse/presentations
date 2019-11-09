package something

type Doer interface {
	Do() error
}

type specialDoer struct{}

func (d specialDoer) Do() error {
	return nil
}

func NewDoer() Doer {
	return specialDoer{}
}
