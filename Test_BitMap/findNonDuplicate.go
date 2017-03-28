/* Calculate count of non-duplicate integers
 * from 250,000,000 positive integers.
 */
package main

import (
	"bitmap"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var l, length, cnt uint64
	var bm *bitmap.BitMap
	length = 1 << 32
	bm = bm.NewBitMap(length, 2)
	for scanner.Scan() {
		l, _ = strconv.ParseUint(scanner.Text(), 10, 64)
		bs := bm.Get(l)
		if bs == "00" {
			bm.Set(l, "01")
			cnt++
		}
		if bs == "01" {
			bm.Set(l, "10")
			cnt--
		}
	}
	fmt.Println(cnt)
}
