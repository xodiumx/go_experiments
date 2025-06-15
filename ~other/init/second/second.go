package second

import "fmt"

func init() {
	fmt.Println("First init in second package")
}

func init() {
	fmt.Println("Second init in second package")
}

func Get() string {
	return "Test"
}
