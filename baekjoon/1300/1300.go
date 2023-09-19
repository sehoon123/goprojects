package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	m := binary_search(n, k)

	//2 * m = (i + j)
	result := lessThan(n, 2*m)

	//fmt.Println(result)

	for i := 2*m + 1; i <= 2*m+2; i++ {
		for j := n; j >= i-j && i-j > 0; j-- {
			if j == i-j {
				result += 1
			} else {
				result += 2
			}
			if result == k {
				fmt.Println((i - j) * j)
				return
			}
		}
	}

}

//// k보다 작은 제곱수중 가장 큰 수를 리턴
//func binary_search(n, k int) int {
//	start := 1
//	end := n
//	for start < end {
//		mid := (start + end) / 2
//		if mid*mid <= k {
//			start = mid + 1
//		} else {
//			end = mid
//		}
//	}
//	return end - 1
//}

func binary_search(n, k int) int {
	start := 1
	end := n
	for start < end {
		mid := (start + end) / 2
		if lessThan(n, 2*mid) <= k {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end - 1
}

// m = i + j 일때 m보다 작은값의 개수를 리턴
func lessThan(n, m int) int {
	return n*(m-n) + (2*n-m)*(m-1)/2
}
