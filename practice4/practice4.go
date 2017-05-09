/*
Reynolds number is calculated using formula (D*v*rho)/mu Where D = Diameter, V= velocity, rho = density mu = viscosity. Write a program that will accept all values in appropriate units (Don't worry about unit conversion) If number is < 2100, display Laminar flow, If its between 2100 and 4000 display 'Transient flow' and if more than '4000', display 'Turbulent Flow' (If, else, then...)
レイノルド数を下記公式を使用して計算する。全ての数字の単位は適切なものが入力されるとする。（単位変換はケアしなくてよい）もし計算結果のレイノルド数が2100未満ならば"Laminar flow"、2100から4000の間であれば"Transient flow"、4000以上であれば"Turbulent flow"と表示する。
D*V*rho/mu, where D=Diameter, V=Velocity, rho=density, mu=viscosity(粘度)
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ReynoldsNumParams struct {
	D   float64
	V   float64
	Rho float64
	Mu  float64
}

func round(f float64, precision int) float64 {
	return math.Floor(f*(math.Pow10(precision))+.5) * math.Pow10(-precision)
}

func Rand(precision int) float64 {
	return round(rand.Float64(), precision)
}

func RandNums(n, precision int) []float64 {
	nums := []float64{}

	for i := 0; i < n; i++ {
		nums = append(nums, Rand(precision))
	}

	return nums
}

func CalcReynoldsNum(d, v, rho, mu float64) (f float64, e error) {
	if mu == 0.0 {
		return 0.0, e
	}

	rNum := (d * v * rho) / mu
	return round(rNum, 2), e
}

func Print(rNum float64) {
	if rNum < 2100.0 {
		fmt.Println("Laminar flow")
	} else if rNum < 4000.0 {
		fmt.Println("Transient flow")
	} else {
		fmt.Println("Turbulent flow")
	}
}

func main() {
	for i := 0; i < 5; i++ {
		nums := RandNums(4, 2)
		params := ReynoldsNumParams{nums[0], nums[1], nums[2], nums[3]}

		fmt.Printf("Reynolds num by %+v => ", params)
		rNum, e := CalcReynoldsNum(params.D, params.V, params.Rho, params.Mu)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(rNum)
			Print(rNum)
		}
	}
}
