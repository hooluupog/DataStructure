package kmp

func Kmp(sstring string,pstring string,next []int) int {
  i, j := 0, 0
	for i < len(sstring) && j < len(pstring) {
		if j == -1 || sstring[i] == pstring[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return i - len(pstring)
}

func Get_next(pstring string) []int {
	next := make([]int, len(pstring))
	next[0] = -1
	for i, j := 0, -1; i < len(pstring)-1; {
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
