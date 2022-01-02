package main

import (
	"fmt"
)

func main() {
	i := 10
	if i%2 == 0 {
		fmt.Println(i, "even")
	} else {
		fmt.Println(i, "odd")
	}

}
