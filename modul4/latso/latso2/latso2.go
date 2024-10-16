package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukkan skor BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat Badan anda adalah: %.0f", beratBadan)
}
