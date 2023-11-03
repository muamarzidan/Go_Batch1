package main 

import "fmt"

func main() {
	var t, ember, volume int
	ember = 0
	fmt.Scan(&t)
	for t != ember {
		fmt.Scan(&volume)
		ember = ember + volume
		fmt.Println(ember == t)
	}
}

// var tangki, ember, volume int
// fmt.Scan(&tangki)
// for tangki != volume {
// 	fmt.Scan(&ember)
// 	volume += ember
// }