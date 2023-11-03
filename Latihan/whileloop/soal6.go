package main 

import "fmt"

func main() {
	var t, ember, volume int;
	ember = 0;
	fmt.Scan(&t);
	
	for t != ember {
		fmt.Scan(&volume)
		ember = ember + volume
		fmt.Println(ember == t)
	}
}
