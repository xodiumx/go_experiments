package main

import (
	"fmt"
	_ "other/init/second"
)

func main() {
	fmt.Print("Test run in mainTwo package")
}
