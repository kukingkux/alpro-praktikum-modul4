package main

import (
	"fmt"
	"math"
)

type koorA struct {
	xA, yA float64
}
type koorB struct {
	xB, yB float64
}
type koorC struct {
	xC, yC float64
}

func main() {
	var (
		AB, BC, CA, panjangsisisegitiga float64
		A                               koorA
		B                               koorB
		C                               koorC
	)

	fmt.Print("Masukkan koordinat A: ")
	fmt.Scanln(&A.xA, &A.yA)
	fmt.Print("Masukkan koordinat B: ")
	fmt.Scanln(&B.xB, &B.yB)
	fmt.Print("Masukkan koordinat C: ")
	fmt.Scanln(&C.xC, &C.yC)
	AB = math.Sqrt(math.Pow(B.xB-A.xA, 2) + math.Pow(B.yB-A.yA, 2))
	BC = math.Sqrt(math.Pow(C.xC-B.xB, 2) + math.Pow(C.yC-B.yB, 2))
	CA = math.Sqrt(math.Pow(A.xA-C.xC, 2) + math.Pow(A.yA-C.yC, 2))

	if AB >= BC && AB >= CA {
		panjangsisisegitiga = AB
	} else if BC >= AB && BC >= CA {
		panjangsisisegitiga = BC
	} else {
		panjangsisisegitiga = CA
	}
	fmt.Printf("%.2f", panjangsisisegitiga)
	// println(AB, BC, CA, panjangsisisegitiga)
}
