package main

import "fmt"

func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) func(d float64) float64 {
	a := acceleration
	v := velocity
	d := displacement
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v*t + d
	}

	return fn

}

func main() {
	var a, v, d, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&d)

	fn := GenDisplaceFn(a, v, d)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fmt.Printf("Displacement: %.2f\n", fn(t))
}
