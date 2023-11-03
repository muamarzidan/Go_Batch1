package main

import main 

func main() {
	var saldo, tf int;
	saldo = 0;

	for tf != 0 {
		saldo = saldo + tf
		fmt.Scan(&tf)
	};
	fmt.Println(saldo);
}