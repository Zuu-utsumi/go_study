/*
Scientific calculator supporting addition, subtraction, multiplication, division, square-root, square, cube, sin, cos, tan, Factorial, inverse, modulus
関数電卓を作る。（足し算、引き算、掛け算、割り算、平方根、３乗根、サイン、コサイン、タンジェント、階乗、反数、剰余）
*/
package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var calcRe = regexp.MustCompile(`([a-zA-Z]+)\((.+)\)`)

func PrintHelp() {
	fmt.Println(`	Commands:
	ADD(x, y)	: x + y
	SUB(x, y)	: x - y
	MUL(x, y)	: x * y
	DIV(x, y)	: x / y
	MOD(x, y)	: x % y
	SQRT(x) 	: √x
	SQUARE(x)	: x ** 2
	CUBE(x) 	: x ** 3
	SIN(x)  	: sin(x)
	COS(x)  	: cos(x)
	TAN(x)  	: tan(x)
	FACT(x) 	: x!
	INV(x)  	: 1 / x`)
}

type Calculator struct {
	vars []float64
	calc func([]float64) float64
}

func (c *Calculator) Calc() (f float64) {
	if c.calc == nil {
		return
	}

	return c.calc(c.vars)
}

func add(vars []float64) float64 {
	return vars[0] + vars[1]
}

func sub(vars []float64) float64 {
	return vars[0] - vars[1]
}

func multiple(vars []float64) float64 {
	return vars[0] * vars[1]
}

func divide(vars []float64) float64 {
	if vars[1] == 0.0 {
		return 0.0
	}

	return vars[0] / vars[1]
}

func mod(vars []float64) float64 {
	return math.Mod(vars[0], vars[1])
}

func sqrt(vars []float64) float64 {
	return math.Sqrt(vars[0])
}

func square(vars []float64) float64 {
	return math.Pow(vars[0], 2)
}

func cube(vars []float64) float64 {
	return math.Pow(vars[0], 3)
}

func sin(vars []float64) float64 {
	return math.Sin(vars[0])
}

func cos(vars []float64) float64 {
	return math.Cos(vars[0])
}

func tan(vars []float64) float64 {
	return math.Tan(vars[0])
}

func factorial(vars []float64) float64 {
	value := vars[0]

	for i := value - 1.0; i >= 1.0; i-- {
		value *= i
	}

	return value
}

func inverse(vars []float64) float64 {
	if vars[0] == 0.0 {
		return 0.0
	}

	return 1.0 / vars[0]
}

// NOTE: dare to implement logic difficult for study
func CalculateStrategy(fname string, vars []float64) (c Calculator) {
	switch strings.ToUpper(fname) {
	case "ADD":
		if len(vars) != 2 {
			fmt.Printf("argument is not enough %v", vars)
			return
		}

		return Calculator{vars, add}
	case "SUB":
		if len(vars) != 2 {
			fmt.Printf("argument is not enough %v", vars)
			return
		}

		return Calculator{vars, sub}
	case "MUL":
		if len(vars) != 2 {
			fmt.Printf("argument is not enough %v", vars)
			return
		}

		return Calculator{vars, multiple}
	case "DIV":
		if len(vars) != 2 {
			fmt.Printf("argument is not enough %v", vars)
			return
		}

		return Calculator{vars, divide}
	case "MOD":
		if len(vars) != 2 {
			fmt.Printf("argument is not enough %v", vars)
			return
		}

		return Calculator{vars, mod}
	case "SQRT":
		return Calculator{vars, sqrt}
	case "SQUARE":
		return Calculator{vars, square}
	case "CUBE":
		return Calculator{vars, cube}
	case "SIN":
		return Calculator{vars, sin}
	case "COS":
		return Calculator{vars, cos}
	case "TAN":
		return Calculator{vars, tan}
	case "FACT":
		return Calculator{vars, factorial}
	case "INV":
		return Calculator{vars, inverse}
	default:
		return
	}
}

func Parse(s string) (c Calculator, e error) {
	m := calcRe.FindStringSubmatch(s)

	if len(m) != 3 {
		// TODO return error
		return
	}

	splitted := strings.Split(m[2], ",")
	vars := []float64{}
	for _, v := range splitted {
		f, e := strconv.ParseFloat(v, -1)

		if e == nil {
			vars = append(vars, f)
		}
	}

	// TODO: validate calculator and return error when problem
	calculator := CalculateStrategy(m[1], vars)

	return calculator, e
}

func main() {
	fmt.Println("Scientific calculator")

	for {
		var s string
		fmt.Print("> ")
		fmt.Scanln(&s)

		if s == "" {
			continue
		}

		if s == "h" || s == "help" {
			PrintHelp()
			continue
		}

		calculator, e := Parse(s)

		if e != nil {
			fmt.Printf("Invalid input %s\n", s)
			continue
		}

		fmt.Printf("= %f\n", calculator.Calc())
	}
}
