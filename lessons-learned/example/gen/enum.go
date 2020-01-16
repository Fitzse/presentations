//go:generate enumer -type=Animal -json -sql
package enum

type Animal int

const (
	KindUnknown Animal = iota
	KindCat
	KindDog
	KindKoala
)
