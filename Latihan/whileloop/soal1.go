package main

import "fmt"

func main() {
    var inputUsername, inputPassword string;
    var i int = 1;
    var login int = -1;

    var username string = "admin";
    var password string = "admin";

    for i > 0 && (inputUsername != username || inputPassword != password) {
        fmt.Scan(&inputUsername)
        fmt.Scan(&inputPassword)
        login++
    };

    fmt.Println(login);
    fmt.Print("Login berhasil");
}
