package main

import "fmt"

func main() {
	var beratKantong1, beratKantong2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kantong pertama: ")
		_, err1 := fmt.Scan(&beratKantong1)
		fmt.Print("Masukan berat belanjaan di kantong kedua: ")
		_, err2 := fmt.Scan(&beratKantong2)

		// Cek jika input tidak valid
		if err1 != nil || err2 != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}

		// Cek jika salah satu berat negatif
		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung total berat
		totalBerat := beratKantong1 + beratKantong2

		// Cek jika total berat lebih dari 150 kg
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat
		selisihBerat := beratKantong1 - beratKantong2
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}

		// Tampilkan apakah sepeda motor oleng atau tidak
		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
