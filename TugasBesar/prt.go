package main
import "fmt"
func main(){
    var noawal, noakhir int
    var berwarna, format string
    var harga, hargaA4, hargaA5, hargaA6 int
    harga = 0
    fmt.Println("pilih (berwarna, tidakberwarna)")
    fmt.Scan(&berwarna)
    if berwarna == "berwarna"{
        hargaA4 = 300
        hargaA5 = 200
        hargaA6 = 100
        fmt.Println("print", berwarna)
        fmt.Println("masukan no awal halaman dan akhir halaman")
        fmt.Scan(&noawal,&noakhir)
        fmt.Println("pilih format kertas (a4, a5, a6)")
        fmt.Scan(&format)
        if format == "a4"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA4 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else if format == "a5"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA5 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else if format == "a6"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA6 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else {
            fmt.Println("ERROR: Masukkan inputan yang benar!")
        }
        
    } else if berwarna == "tidakberwarna" {
        hargaA4 = 250
        hargaA5 = 150
        hargaA6 = 100
        fmt.Println("print tidak berwarna")
        fmt.Println("masukan no awal halaman dan akhir halaman")
        fmt.Scan(&noawal,&noakhir)
        fmt.Println("pilih format kertas (a4, a5, a6)")
        fmt.Scan(&format)
        if format == "a4"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA4 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else if format == "a5"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA5 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else if format == "a6"{
            for i := noawal; i <=noakhir; i++{
                harga = hargaA6 * i
            }
            fmt.Println("jumlah halaman", noakhir)
            fmt.Println("format",format,"berwarna")
            fmt.Println("harga", harga)
        } else {
            fmt.Println("ERROR: Masukkan inputan yang benar!")
        }
    } else {
        fmt.Println("ERROR: Masukkan inputan yang benar!")
    }
}