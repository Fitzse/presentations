func DoStuff(s []string) {
	m := make(map[string]int, len(s))
	for _, str := range s {
		m[str] = something(str)
	}
}