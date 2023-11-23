package main 

import "fmt"

func main () {
	var n, m int;
	fmt.Scan(&n, &m)

	for n <= m {
		fmt.Println(n)
		n = n + 1
	}
	for i := n; i <= m; i++ {
		fmt.Println(i)
	}
}

program maxn
kamus
	n, m : integer
algoritma
	input(n, m)
	while n <= m do
		output(n)
		n <- n + 1
	end while
end program