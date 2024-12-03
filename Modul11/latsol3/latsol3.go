package main

import "fmt"

func main() {
    var bilangan int
    fmt.Print("Masukkan sebuah bilangan: ")
    fmt.Scan(&bilangan)

    switch {
    case bilangan%10 == 0:
        hasil := bilangan / 10
        fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
    case bilangan%5 == 0 && bilangan > 5:
        hasil := bilangan * bilangan
        fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d^2 = %d\n", bilangan, hasil)
    case bilangan%2 == 0:
        hasil := bilangan * (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
    case bilangan%2 != 0:
        hasil := bilangan + (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
    default:
        fmt.Println("Nilai tidak valid.")
    }
}
