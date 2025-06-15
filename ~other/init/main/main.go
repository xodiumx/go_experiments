package main

import (
	"fmt"
)

func init() {
	fmt.Println("Third init in main package")
}

func init() {
	fmt.Println("Fourth init in main package")
}

//func main() {
//	sec.Get()
//}
