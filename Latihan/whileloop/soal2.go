package main

import "fmt"


func main() {
    var hasil, tf int;
    hasil = 0;

    for {
        fmt.Scan(&tf)
        if tf == 0{
			break
		}
        hasil = hasil + tf
    };
    fmt.Println(hasil);
};
