package test

import "fmt"

var(
	Animal int
	Plant int
	dont int

)

func init()  {
	Animal = 3
	Plant = 99
	dont = -10
	fmt.Println("init test")
}