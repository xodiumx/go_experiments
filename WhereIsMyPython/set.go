package main

import "fmt"

type Set map[string]bool

func (s Set) Add(key string) {
	s[key] = true
}

func (s Set) Contains(key string) bool {
	return s[key]
}

func (s Set) Remove(key string) {
	delete(s, key)
}

func main() {
	my_set := make(Set)

	my_set.Add("a")
	my_set.Add("b")
	my_set.Add("c")

	fmt.Println(my_set.Contains("a")) // true
	fmt.Println(my_set.Contains("b")) // true
	fmt.Println(my_set.Contains("W")) // false
}
