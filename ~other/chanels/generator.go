package main

func isPalindromeNum(x int, base int) bool {
	if x%base == 0 {
		return x == 0
	}
	reversed := 0
	original := x
	for x > 0 {
		reversed = reversed*base + x%base
		x /= base
	}
	return original == reversed
}

func decimalPalindromeGeneratorSlow() func() int {
	i := 1
	return func() int {
		for {
			if isPalindromeNum(i, 10) {
				val := i
				i++
				return val
			}
			i++
		}
	}
}

//func main() {
//	n := 20
//	k := 5
//	var total int
//	var results []int
//	gen := decimalPalindromeGeneratorSlow()
//	for len(results) != n {
//		for {
//			val := gen()
//			if isPalindromeNum(val, k) {
//				total += val
//			}
//
//		}
//	}
//	fmt.Println(total)
//}
