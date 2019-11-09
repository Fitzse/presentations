package somethingelse

type Doer interface {
	Do() error
}

func Handle(d Doer) error {
	return d.Do()
}
