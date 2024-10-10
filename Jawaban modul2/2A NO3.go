package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari float64

	// Input jari-jari dari pengguna
	fmt.Print("Jejari = ")
	fmt.Scan(&jejari)

	// Menghitung volume dan luas kulit bola
	volumeBola := (4.0 / 3.0) * math.Pi * math.Pow(jejari, 3)
	luasBola := 4 * math.Pi * math.Pow(jejari, 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volumeBola, luasBola)
}
