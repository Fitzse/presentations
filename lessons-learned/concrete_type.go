package something

type Thing struct{}

func (t Thing) Do() error {
	return nil
}

func NewThing() Thing {
	return Thing{}
}
