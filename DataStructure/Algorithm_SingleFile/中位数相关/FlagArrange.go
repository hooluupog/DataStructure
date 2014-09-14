/* 划分算法的应用 */
/* 荷兰国旗问题：设有一个仅有红、白、蓝三种颜色组成的条块组成条块序列，*/
/* 编写一个复杂度O(n)的算法，使得这些条块按红、白、蓝顺序排好 */
/* 即组成荷兰国旗图案  */
package main

import "fmt"

const (
	RED = iota
	WHITE
	BLUE
)

func FlagArrange(s string) string {
	b := []byte(s)
	i, j, k := 0, 0, len(b)-1
	for j <= k {
		switch b[j] - '0' {
		case RED:
			b[i], b[j] = b[j], b[i]
			i++
			j++
		case WHITE:
			j++
		case BLUE:
			b[j], b[k] = b[k], b[j]
			k-- // 这里没有j++语句，防止交换后b[j]仍为蓝色的情况
		}
	}
	return string(b)
}
func main() {
	var s string
	fmt.Scan(&s)
	res := FlagArrange(s)
	fmt.Println(res)
}
