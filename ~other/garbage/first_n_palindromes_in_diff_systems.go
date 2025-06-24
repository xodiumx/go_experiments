package main

import (
	"strconv"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func intPow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func decimalPalindromeGenerator() func() int {
	queue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := 1

	return func() int {
		if len(queue) > 0 {
			val := queue[0]
			queue = queue[1:]
			return val
		}

		// начинаем строить палиндромы > 9
		length++

		start := intPow(10, (length-1)/2)
		end := intPow(10, (length+1)/2)

		for i := start; i < end; i++ {
			s := strconv.Itoa(i)
			var palStr string
			if length%2 == 0 {
				palStr = s + reverseString(s)
			} else {
				palStr = s + reverseString(s[:len(s)-1])
			}
			val, _ := strconv.Atoi(palStr)
			queue = append(queue, val)
		}

		// повторный вызов достанет первый сгенерированный
		val := queue[0]
		queue = queue[1:]
		return val
	}
}

//func main() {
//	n := 20
//	k := 5
//	var total int
//	var counter int
//	gen := decimalPalindromeGenerator()
//	for counter != n {
//		num := gen()
//		val := strconv.FormatInt(int64(num), k)
//		if isPalindrome(val) {
//			fmt.Printf("Number in 10 system: %v Number in k system: %v Counter: %v\n", num, val, counter)
//			total += num
//			counter++
//		}
//	}
//	fmt.Printf("Total: %v", total)
//}
