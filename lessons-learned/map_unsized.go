func DoStuff(s []string) {
	m := map[string]int{}
	for _, str := range s {
		m[str] = something(str)
	}
}