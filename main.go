package main

import "fmt"

func main() {
	var x int 
	var a, b, c string

	level1Satu := true
	level1Dua := true
	level1Tiga := true
	
    fmt.Println("|---------->SAMBUNG LIRIK LAGU<----------|")
    fmt.Print("Pilih tingkat kesulitan(1,2,3) : ")
    fmt.Scan(&x)


    for i := 1; i <= x; i++ {
        if i == 1 && x == 1{
            for j := 1; j <= 3; j++ {
                if j == 1{
                    fmt.Println("satu ditambah satu")
                    fmt.Scan(&a)
					level1Satu = a == "dua"
					fmt.Println(level1Satu)
                } else if j == 2{
					fmt.Println("dua ditambah dua")
					fmt.Scan(&b)
					level1Dua = b == "empat"
					fmt.Println(level1Dua)
				} else if j == 3{
					fmt.Println("tiga ditambah tiga")
					fmt.Scan(&c)
					level1Tiga = c == "enam"
					fmt.Println(level1Tiga)
				}
            }
        }
        if i == 2 && x == 2 {
            fmt.Println("knt")
        }
        if i == 3 && x == 3 {
            fmt.Println("benar")
        }
	}
	// make a program to see if an integer input from the user has digits that are in decreasing order or not.
	// the input consists of positive integers, the first number is N, while the second or next number is a number consisting of N digits.
	// the output consists of a boolean stating that the given second number is in decreasing order or not.
	
	// example 
	// input 
	// 4 1259
	// 5 98431
	
	// output
	// false
	// true
	
	// Note:
	// must use for loop/while loop/repeat until
	// must not use branching, additional libraries etc.

		// var n int;
		// fmt.Scan(&n);

		// n = n / 10;
		// n = n % 10;

		// fmt.Printf("%d\n", n)

	



	// var n, m, hasil int

	// fmt.Scan(&n, m)
	// hasil = 0

	// for i := 0; i < n; i++ {
	// 	hasil = n + m
	// }

	// fmt.Println(hasil)

		

	// var x, y, z int
	// var t bool
	// fmt.Scan(&x)
	// t = true

	// for x > 9 && t {
	// 	z = x % 10
	// 	y = (x % 100) / 10
	// 	t = z-y == 1 || y-z == 1
	// 	x = x / 100
	// }
	// fmt.Println(t)



	// var bil, bil2, bilSeb int
	// var checkBil, checkBilSeb bool

	// fmt.Scan(&bil)

	// checkBil = bil > 0
	// for checkBil {
	// 	bil2 = bil % 10
	// 	bil = bil / 10
	// 	bilSeb = bil % 10

	// 	checkBilSeb = bil2-bilSeb == 1 || bilSeb-bil2 == 1 || bil2-bilSeb == -1 || bilSeb-bil2 == -1
	// 	checkBil = checkBilSeb
	// }
	// fmt.Println(checkBil)
}

// func main () {
// 	var bil,bilAwal,BilKedua int;
// 	fmt.Scan(&bil);

// 	bilAwal = bil / 10;
// 	BilKedua = bil % 10;

// 	fmt.Printf("%d %d\n", bilAwal, BilKedua);
// }

// func main() {
// 	var bil float64
// 	var bilBulat int64
// 	fmt.Scan(&bil)

// 	bilBulat = int64(bil)
// 	bilDes := bil - float64(bilBulat)

// 	fmt.Printf("%f %d",bilDes, bilBulat)
// }

// func main () {
// 	var bil, i int
// 	fmt.Scan(&bil)
// 	for i = bil; i > 0; i-- {
// 		fmt.Print(i, "x")
// 	}
// }

// func main () {
// 	var input, bil1, bil2, bil3, bil4, bil5 int
// 	fmt.Scan(&input)

// 	bil1 = input / 10000
// 	bil2 = input / 1000 % 10
// 	bil3 = input / 100 % 10
// 	bil4 = input / 10 % 10
// 	bil5 = input % 10

// 	fmt.Printf("%d %d %d %d %d\n", bil5, bil4, bil3, bil2, bil1)
// }

// func main () {
// 	var xi, total int;
// 	total  = 0;
// 	fmt.Scan(&xi)
// 	for xi % 2 == 0 {
// 		fmt.Scan(xi)
// 		total = total + xi
// 	}
// 	fmt.Println(total)
// }

// func main () {
// 	var semester int;
// 	var ipk float64;
// 	var publiskasi, lulus bool;

// 	fmt.Scan(&semester, &ipk, &publiskasi		)

// 	lulus = semester <= 8 && semester >=7 && ipk >= 3.50 && publiskasi

// 	fmt.Println(lulus)
// }

// func main() {
// 	var id, d, m, y int

// 	fmt.Scan(&id)

// 	d = (id / 100000) / 100 % 100
// 	// m = (id / 100000) % 100
// 	// y = (id % 100000) / 10

// 	fmt.Printf("saya lahir pada tanggal %d bulan ke %d dan tahun ke %d\n", d, m, y)
// }
