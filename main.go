package main

import (
	"fmt"

	"github.com/34vsk/golimp/kopas"
	"github.com/34vsk/golimp/pizza"
	"github.com/34vsk/golimp/prime"
)

func main() {
	fmt.Println("Raksturīgās kopas")
	kopas.Solve()
	fmt.Println("Divu pirmskaitļu reizinājums")
	prime.Solve()
	fmt.Println("Picērija")
	pizza.Solve()
}
