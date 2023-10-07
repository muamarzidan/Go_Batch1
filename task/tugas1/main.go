package main

import (
	"fmt"
	// "math"
	// "strconv"
	// "strings"
)

func main() {
	fmt.Println("NAMA : Muamar Zidan Tri Antoro")
	fmt.Println("NIM : 103012300381")
	fmt.Println("KELAS : IF-47-11")
}

//soal no 1
// func main() {
// 	var jariJariAwal string
// 	const phi float64 = 3.14

// 	fmt.Print("Jari-jari lingkaran : ")
// 	fmt.Scan(&jariJariAwal)

// 	jariJari, err := strconv.ParseFloat(jariJariAwal, 64)
// 	if err != nil {
// 		fmt.Println("Pastikan inputan yang dimasukan adalah angka angka")
// 		return
// 	}

// 	r := jariJari

// 	luasLingkaran := phi * r * r
// 	luasLingkaran = math.Round(luasLingkaran*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

// 	fmt.Println("Jadi, dengan jari-jari", jariJari, "Luas lingkaran adalah:", luasLingkaran)
// }

//soal no 2
// func main() {
//     const phi float64 = 3.14
//     var luasLingkaran float64

//     fmt.Print("diketahui luas lingkaran : ")
//     fmt.Scan(&luasLingkaran)

//     var diameter float64 = 2 * math.Sqrt(luasLingkaran/phi)
//     diameter = math.Round(diameter*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

//     fmt.Println("Jadi, dengan luas lingkaran", luasLingkaran, ", diameter lingkarannya adalah", diameter)
// }

//soal no 3
// func main() {
// 	var celcius, reamur float64
// 	fmt.Print("Masukan angka untuk suhu celcius : ")
// 	fmt.Scan(&celcius)

// 	reamur = 4 * celcius / 5
// 	reamur = math.Round(reamur*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

// 	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan reamur adalah :", reamur, "derajat reamur")
// }

//soal no 4
// func main() {
// 	var panjang, lebar float64

// 	fmt.Print("Panjang persegi panjang : ")
// 	fmt.Scan(&panjang)

// 	fmt.Print("Lebar persegi panjang : ")
// 	fmt.Scan(&lebar)

// 	var kelilingPersegiPanjang float64  = 2 * (panjang + lebar)
// 	kelilingPersegiPanjang = math.Round(kelilingPersegiPanjang * 1) / 1 //handle if the input is not a bilangan bulat

// 	fmt.Println("Jadi, dari panjang", panjang, "dan lebar", lebar, ", maka keliling persegi panjang adalah :", kelilingPersegiPanjang)
// }

//soal no 5
// func main () {
// 	var celcius, fahrenheit float64
// 	fmt.Print("Mausukan angka untuk suhu celcius : ")
// 	fmt.Scan(&celcius)

// 	fahrenheit = (9 * celcius / 5) + 32
// 	fahrenheit = math.Round(fahrenheit * 100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

// 	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan fahrenhait adalah :", fahrenheit, "derajat farhenhait")
// }

//soal no 6&7 sama
// func main () {
// 	var celcius, kelvin float64
// 	fmt.Print("Masukan suhu dalam celcius : ")
// 	fmt.Scan(&celcius)

// 	kelvin = celcius + 273.15
// 	kelvin = math.Round(kelvin * 100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

// 	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan kelvin adalah :", kelvin, "derajat kelvin")
// }

//soal no 8
// func main() {
// 	var jariJari, luasPermukaan float64
//     const phi float64 = 3.14

//     fmt.Print("Jari-jari bola : ")
//     fmt.Scan(&jariJari)

//     r := jariJari

//     luasPermukaan = 4 * phi * r * r
//     luasPermukaan = math.Round(luasPermukaan*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

//     fmt.Println("Luas permukaan bola : ", luasPermukaan)
// }

//soal no 9
// func main() {
// 	var bilangan int
// 	var Hasil bool

// 	fmt.Print("Bilangan bilangan bulat : ")
// 	fmt.Scan(&bilangan)

// 	if bilangan%2 != 0 {
// 		Hasil = true
// 		fmt.Println(Hasil)
// 	} else  {
// 		Hasil = false
// 		fmt.Println(Hasil)
// 	}
// }

//soal no 10
// func main() {
//     var bilNegatif int

//     fmt.Print("Masukkan bilangan bulat : ")
//     fmt.Scan(&bilNegatif)

//     if bilNegatif < 0 {
//         fmt.Println("true")
//     } else {
//         fmt.Println("false")
//     }
// }

//soal no 11
// func main() {
//     var bilPositif int
// 	var Hasil bool

//     fmt.Print("Masukkan bilangan bulat : ")
//     fmt.Scan(&bilPositif)

//     if bilPositif > 0 {
// 		Hasil = true
// 		fmt.Println(Hasil)
//     } else {
// 		Hasil = false
// 		fmt.Println(Hasil)
//     }
// }

//soal no 12
// func main() {
// 	var hurufVokal string
// 	var Hasil bool

// 	fmt.Print("Masukkan sebuah huruf : ")
// 	fmt.Scan(&hurufVokal)
// 	hurufVokal = strings.ToLower(hurufVokal)

// 	if hurufVokal == "a" || hurufVokal == "i" || hurufVokal == "u" || hurufVokal == "e" || hurufVokal == "o" {
// 		Hasil = true
// 		fmt.Println(Hasil)
// 	} else {
// 		Hasil = false
// 		fmt.Println(Hasil)
// 	}
// }

//soal no 13&14 sama
// func main() {
// 	var hurufKonsonan string
// 	var Hasil bool

// 	fmt.Print("Masukkan sebuah huruf : ")
// 	fmt.Scan(&hurufKonsonan)
// 	hurufKonsonan = strings.ToLower(hurufKonsonan)

// 	if hurufKonsonan != "a" && hurufKonsonan != "i" && hurufKonsonan != "u" && hurufKonsonan != "e" && hurufKonsonan != "o" {
// 		Hasil = true
// 		fmt.Println(Hasil)
// 	} else {
// 		Hasil = false
// 		fmt.Println(Hasil)
// 	}
// }

// soal no 15
// func main() {
// 	var bilangan int
// 	var Hasil bool

// 	fmt.Print("Masukkan sebuah bilangan bulat : ")
// 	fmt.Scan(&bilangan)

// 	if bilangan == 0 {
// 		Hasil = true
// 		fmt.Println(Hasil)
// 	} else {
// 		Hasil = false
// 		fmt.Println(Hasil)
// 	}
// }

//soal no 16
// func main() {
// 	var namaSaya string

// 	fmt.Print("Masukkan nama : ")
// 	fmt.Scan(&namaSaya)

// 	fmt.Println("Hai, " + namaSaya)
// }

//soal no 17
// func main() {
//     var karakter byte
//     fmt.Print("Masukkan sebuah karakter : ")
// 	fmt.Scanf("%c", &karakter)

//     fmt.Printf("Karakter yang dimasukkan adalah : %c\n", karakter)
// }
