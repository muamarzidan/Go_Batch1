package main 

import "fmt"

func main () {
	var n, faktorial int
	fmt.Scan(&n)

	faktorial = 1

	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Println(faktorial)
}