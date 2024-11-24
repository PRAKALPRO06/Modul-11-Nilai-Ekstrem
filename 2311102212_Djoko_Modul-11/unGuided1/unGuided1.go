package main

import (
	"fmt"
)

type Rabbit struct {
	weight int
}

type arrRabbit [1000]Rabbit

func main() {
	var n2311102212 int
	var rabbits arrRabbit

	fmt.Print("Banyaknya anak kelinci: ")
	fmt.Scan(&n2311102212)

	for i := 0; i < n2311102212; i++ {
		fmt.Printf("Anak kelinci ke-%d, berat: ", i+1)
		fmt.Scan(&rabbits[i].weight)
	}

	maks := beratTerbesar(rabbits, n2311102212)
	min := beratTerkecil(rabbits, n2311102212)

	fmt.Println("Berat Terbesar : ", maks)
	fmt.Println("Berat Terkecil : ", min)
}

func beratTerbesar(T arrRabbit, n int) int {
	maks := T[0].weight 
	for i := 1; i < n; i++ {
		if T[i].weight > maks { 
			maks = T[i].weight
		}
	}
	return maks
}

func beratTerkecil(T arrRabbit, n int) int {
	min := T[0].weight
	for i := 1; i < n; i++ {
		if T[i].weight < min {
			min = T[i].weight
		}
	}
	return min
}
