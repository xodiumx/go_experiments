package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   int // или string, в зависимости от типа ключа
	Value int
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, val := range nums {
		hashMap[val]++
	}
	var keyValues []KeyValue
	for k, v := range hashMap {
		keyValues = append(keyValues, KeyValue{k, v})
	}
	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Value > keyValues[j].Value
	})

	var result []int
	for _, val := range keyValues {
		if len(result) == k {
			return result
		}
		result = append(result, val.Key)
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 4, 4, 4, 9, 10}
	fmt.Print(topKFrequent(nums, 3))
}
