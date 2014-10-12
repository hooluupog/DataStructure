// 归并思想的应用 
// codeforces 414C - Mashmokh and Reverse Operation 
// 题意：给定2^n长的序列，对序列进行m次查询，对于每个询问q:
//  1) 先把序列分成2^(n-q)段: [0, 2^q-1],[2^q,2^(q+1)-1],......,每段长2^q，(0<= q <= n);
//  2) 再把每段都翻转一下;
//  3) 输出此时序列的逆序数。
// 解题思路：
// 2^n个数，可以利用树状数组，比如2,1,4,3：
//
//                                ［2,1,4,3］                  level 1
//                                /         \
//                             [2,1]       [4,3]               level 2
//                            /     \      /   \
//                           [2]    [1]   [4]  [3]             level 3
//
// 把长为2^n的序列分层得到n层，对于每一层[l, r]，逆序数对分为以下三种情况：
// 1)两个数都在[l,mid]内;
// 2)两个数字都在[mid+1,r]内;
// 3)一个数在[l,mid]，一个数在[mid+1,r]内;
// [l,r]的逆序数 = [l,mid]的逆序数 + [mid+1,r]的逆序数 + [l,mid],[mid+1,r]之间的逆序数.（递推式）

// 当区间内只有2个数时，则只会出现第3)种情况，即逆序数对分布在两个子区间内.             (初始值)
// T(l,r,level) = T(l,mid) + T(mid+1,r) + Between(level).
// 为简便起见，考虑理想化的情形：设[l,mid]和[mid+1,r]长度均为2^n-1,每次分割区间都会长度减半,则：
// T(2^n,1) = T1(2^(n-1),1) + T2(2^(n-1),1) + Between1(1)
//          = T1(2^(n-2),2) + T2(2^(n-2),2) + T3(2^(n-2),2) + T4(2^(n-2),2) + Between1(1) + Between2(2) + Between3(2)
//          = T1(1,n) + T2(1,n) + ... + T2^n(1,0) + Between1(1) + Between2(2) + ... + Between(2^n-1)(n)
// 显然，T1(1,n) = T2(1,n) = ... T2^n(1,0) = 0(区间只有一个元素，逆序数为0).
// 则逆序数 = T(2^n,1) = Between1(1) + Between2(2) + ... + Between(2^n-1)(n).
// 上式说明1)和2)两种情况的逆序数可以由3)得出，即同在一个区间内的逆序数就是其对应的下一层的两个子区间之间的逆序数。
// 例如[2,1,4,3]的逆序数就是它的下一层[2,1],[4,3]两个子区间之间的逆序数。
// 累加每层的区间之间的逆序数即可得到整个序列的逆序数。
// 翻转分析：
// 对于一个序列[l,r]翻转
// 1)将该序列的[l,mid], [mid+1,r]交换 
// 2)[l,mid/2],[mid/2+1,mid]交换 + [mid,mid+mid/2],[mid+mid/2+1,r]交换
// 3)如此递归下去
// 例如：12345678-> 87654321,翻转过程：
//         1234|5678 -> 5678|1234 -> 7856|3412 -> 8765|4321
// 假如序列分成2^k段，翻转第k+1层子序列对于level k+1层以上的逆序数是不会改变的，但是level k~level n的逆序数对会翻转，
// 我们只需要知道level k～level n各个区间翻转后的逆序数就可以了。 
// 具体过程：
// 对数组进行一次归并排序，在排序的过程中，每当发生归并时(即将两个有序的子序列归并为一个有序的序列之前)：
// 统计[l,mid]和[mid+1,r]之间的逆序数;
// 统计翻转后的两个子区间[mid+1,r]和[l,mid]之间的逆序数。
// 用一个二维数组sum[N+1][2]保存上面的逆序数,
// sum[level][0]:表示当前层[l,mid,mid+1,r]区间的逆序数
// sum[level][1]:表示当前层[l,mid,mid+1,r]区间翻转后[mid+1,r,l,mid]的逆序数。
// 递归排序完成之后，根据每次的查询q的取值，翻转对应的子序列,即先交换sum[level][0]与sum[level][1],再依次累加每层的逆序数,
// answer += sum[level][0],最终可以得到整个序列的逆序数。


package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const N = 20

var (
	a   []int64 = make([]int64, (1<<N)+1)
	aux []int64 = make([]int64, (1<<N)+1)  // 归并排序的辅助数组
	sum [N + 1][2]int64                    // 每一层的逆序数
)

func Reverse(a []int64, low int64, high int64) { //逆置序列
	for i := low; i <= (low+high)/2; i++ {
		a[i], a[high-i+low] = a[high-i+low], a[i]
	}
}

func LowerBound(a []int64, low, high int64, value int64) int64 {
// 查找升序序列中value的插入位置,即比value小最后一个a[i]元素的位置
	l, h := low, high
	pos := low - 1
	for l <= h {
		mid := (l + h) >> 1
		if value > a[mid] {
			l = mid + 1
			pos = mid
		} else {
			h = mid - 1
		}
	}
	return pos
}

func Merge(aux []int64, a []int64, low, mid, high int64) { // 合并两个有序子序列
	for i, j, k := low, mid+1, low; k <= high; k++ {
		if j > high || i <= mid && aux[i] <= aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}
func Inversions(aux []int64, a []int64, low int64, high int64, deep int64) {// 统计逆序数
	if low >= high {
		return
	}
	mid := (low + high) >> 1
	Inversions(a, aux, low, mid, deep-1)
	Inversions(a, aux, mid+1, high, deep-1)
	count := int64(0)
	for i, j := low, mid; i <= mid; i++ { // [low,mid]和[mid+1,high]之间的逆序数
		j = LowerBound(aux, j+1, high, aux[i])
		count = count + j - mid
	}
	sum[deep][0] = sum[deep][0] + count
	// 翻转后的逆序数统计
	count = 0
	for i, j := mid+1, low-1; i <= high; i++ { // [mid+1,high]和[low,mid]之间的逆序数
		j = LowerBound(aux, j+1, mid, aux[i])
		count = count + j - low + 1
	}
	sum[deep][1] = sum[deep][1] + count
	if aux[mid] <= aux[mid+1] { // 已经有序，直接复制到a当中
		copy(a[low:high+1], aux[low:high+1])
		return
	}
	Merge(aux, a, low, mid, high)
}

func main() {
	var n, m, que int64
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscan(r, &n)
		if err == io.EOF {
			break
		}
		for i := 0; i < len(sum); i++ {
			sum[i][0] = int64(0)
			sum[i][1] = int64(0)
		}
		length := int64(math.Pow(2, float64(n)))
		a[0] = length
		for i := int64(1); i <= length; i++ {
			fmt.Fscan(r, &a[i])
		}
		copy(aux, a)
		Inversions(aux, a, 1, a[0], n)
		fmt.Fscan(r, &m)
		for i := int64(1); i <= m; i++ {
			fmt.Fscan(r, &que)
			ans := int64(0)
			for i := que; i >= 1; i-- { //根据que的值先翻转序列
				sum[i][0], sum[i][1] = sum[i][1], sum[i][0]
			}
			//将翻转后的结果累加
			for i := int64(0); i <= n; i++ {
				ans = ans + sum[i][0]
			}
			fmt.Fprintln(w, ans)
		}
	}
	w.Flush()
}
