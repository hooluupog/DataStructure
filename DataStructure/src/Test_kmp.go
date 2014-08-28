package main
import (
    "fmt"
    "kmp"
)

func main() {
        sstring := "acabaabaabcacaabc"
	pstring := "abaabcac"
        next := make([]int,len(pstring))
	next = kmp.Get_next(pstring)
	fmt.Printf("The matching postion is : %d\n",kmp.Kmp(sstring, pstring,next))
}
