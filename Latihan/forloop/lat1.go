package main  

import "fmt"

func main () {
	var n int;	
	var text string;

	fmt.Scan(&n, &text)

	for i := 0; i < n; i++ {
		fmt.Println(text)
	}
}