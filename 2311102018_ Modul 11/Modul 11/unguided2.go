//2311102018-Aryo Tegar Sukarno
package main

import "fmt"
const maxCapacity = 1000
const pricePerKg = 50000 

func calculateTotalPrice(weights [maxCapacity]float64, n int) float64 {
	var totalPrice float64 = 0
	for i := 0; i < n; i++ {
		totalPrice += weights[i] * pricePerKg
	}
	return totalPrice
}
func main() {
	var weights [maxCapacity]float64
	var n int
	fmt.Print("Masukkan jumlah ikan yang akan dijual (maks 1000): ")
	fmt.Scan(&n)
	if n < 1 || n > maxCapacity {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000.")
		return
	}
	fmt.Println("Masukkan berat ikan (dalam kilogram):")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
		if weights[i] <= 0 {
			fmt.Println("Berat ikan harus lebih dari 0.")
			i--
		}
	}
	totalPrice := calculateTotalPrice(weights, n)
	fmt.Println("\nDaftar berat ikan dan harga jual:")
	for i := 0; i < n; i++ {
		fmt.Printf("Ikan ke-%d: Berat = %.2f kg, Harga = %.2f IDR\n", i+1, weights[i], weights[i]*pricePerKg)
	}
	fmt.Printf("\nTotal harga jual semua ikan: %.2f IDR\n", totalPrice)
}
