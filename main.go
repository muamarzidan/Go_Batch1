package main

import "fmt"

// import "fmt"


func main () {
	var xi, total int;
	total  = 0;
	fmt.Scan(&xi)
	for xi % 2 == 0 {
		fmt.Scan(xi)
		total = total + xi
	}
	fmt.Println(total)
}

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
//MENCETAK STRING
// func main() {
// 	fmt.Printf("Hello, %s! Your age is %d\n", "John", 30)
// }

// func main() {
// 	fmt.Print("Hello, World!2")
// }

// func main() {
// 	fmt.Printf("Hello, World!3")
// }

// func main() {
// 	fmt.Printf("%s\n", "Hello, World!4")
// }

// func main () {
// 	var a string
// 	a = "Hello, World!5"
// 	fmt.Println(a)
// }

//MENCETAK INTEGER
// func main() {
// 	var a int
// 	a = 10
// 	fmt.Println(a)
// }

// func main () {
// 	fmt.Println(3.13)
// }

//MENCETAK CHAR
// func main () {
// 	fmt.Printf("%c\n" , 'A')
// }
