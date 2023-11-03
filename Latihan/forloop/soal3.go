package main

import "fmt"

func main() {
	var jumlahsiswa, waktuHari, totalWaktu int
	fmt.Scan(&jumlahsiswa)

	for i := 0; i < jumlahsiswa; i++ {
		fmt.Scan(&waktuHari)
		totalWaktu = totalWaktu + waktuHari
	}

	rata_rata := float64(totalWaktu) / float64(jumlahsiswa)
	fmt.Printf("%.3f\n", rata_rata) // to show how many number desimal 0 nya, use the point after % nya gabung angka to calculate
}
