package util

func FilterStringSlice(lines []string, f func(l string) bool) []string {
	r := lines[:0]
	for _, l := range lines {
		if f(l) {
			r = append(r, l)
		}
	}
	return r
}
