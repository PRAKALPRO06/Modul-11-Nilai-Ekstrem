package main
import (
	"fmt"
)

func main() {
	var jumlahAnak int
	var BobotBeratAnak [1000] float64
	var beratPalingKecil, beratPalingGede float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahAnak)

	if jumlahAnak < 1 || jumlahAnak > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < jumlahAnak; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&BobotBeratAnak[i])
	}

	beratPalingKecil = BobotBeratAnak[0]
	beratPalingGede = BobotBeratAnak[0]

	for i := 1; i < jumlahAnak; i++ {
		if BobotBeratAnak[i] < beratPalingKecil {
			beratPalingKecil = BobotBeratAnak[i]
		}
		if BobotBeratAnak[i] > beratPalingGede {
			beratPalingGede = BobotBeratAnak[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", beratPalingKecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratPalingGede)
}