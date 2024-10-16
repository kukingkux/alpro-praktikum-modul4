package main

import "fmt"

func main() {
	var (
		totalbelanja, diskon            int
		hargasetelahdiskon, hargadiskon float64
	)
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalbelanja)
	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&diskon)
	hargadiskon = float64(totalbelanja) * (float64(diskon) / 100)
	hargasetelahdiskon = float64(totalbelanja) - hargadiskon
	fmt.Printf("%.0f", hargasetelahdiskon)
}
