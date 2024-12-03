package main

import "fmt"

func main() {
	var tipe string
	var jam, tarif int

	fmt.Scan(&tipe, &jam)
	

	switch {
	case tipe == "Motor" && jam == 1 && jam < 2:
		tarif = 2000
	case tipe == "Motor" && jam >= 2:
		tarif = jam * 2000
	case tipe == "Mobil" && jam == 1 && jam < 2:
		tarif = 5000
	case tipe == "Mobil" && jam >= 2:
		tarif = jam * 5000
	case tipe == "Truk" && jam == 1 && jam <= 2:
		tarif = 8000
	case tipe == "Truk" && jam >=	 2:
		tarif = jam * 8000
	}
	fmt.Print("Rp ", tarif)
}