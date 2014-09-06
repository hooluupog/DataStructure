package kmp

func Kmp(sstring string, pstring string, next []int) []int {
	i, j := 0, 0
	s := []int{} // sucess paterning posistions.
	for i < len(sstring) && len(pstring) <= len(sstring) {
		if j == -1 || sstring[i] == pstring[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == len(pstring) {
			s = append(s, i-len(pstring))
			j = next[j] // next round paterning start position.
		}
	}
	return s
}

func Get_next(pstring string) []int {
	next := make([]int, len(pstring)+1)
	next[0] = -1
	for i, j := 0, -1; i < len(pstring); {
		if j == -1 || pstring[i] == pstring[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}