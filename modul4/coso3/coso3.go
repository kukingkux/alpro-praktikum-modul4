package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukkan Berat Badan (kg): ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukkan Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
