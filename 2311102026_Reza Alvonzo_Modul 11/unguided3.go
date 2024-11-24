package main
import(
	"fmt"
)
type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, x int,bMin, bMax *float64){
  *bMin = arrBerat[0]
  *bMax = arrBerat[0]
  for i := 1; i < x; i++ {
    if arrBerat[i] > *bMax {
      *bMax = arrBerat[i]
    }
    if arrBerat[i] < *bMin {
      *bMin = arrBerat[i]
    }
  }
}

func rerata(arrBerat arrBalita, x int) float64{
  sum:= 0.0
  for i := 0; i < x; i++ {
    sum += arrBerat[i]
  }
  return sum/ (float64(x))
}


func main(){
  var x int
  var arrBalita[100] float64
  var bMin,bMax float64
  
  fmt.Print("Masukkan banyak data berat balita: ")
  fmt.Scan(&x)

  for i := 0; i < x; i++ {
    fmt.Printf("Masukkan berat balita ke-%d: ",i+1)
    fmt.Scan(&arrBalita[i])
  }
  fmt.Println()

  hitungMinMax(arrBalita,x,&bMin,&bMax)
  avg := rerata(arrBalita,x)

  fmt.Printf("Berat balita maksimal: %.2f kg\n", bMax)
  fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
  fmt.Printf("Rerata berat balita: %.2f kg", avg )
}

//Reza Alvonzo IF 11 06 2311102026//