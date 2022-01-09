package main

import (
	"fmt"

	"github.com/doublegrey/olimp/kopas"
	"github.com/doublegrey/olimp/pizza"
	"github.com/doublegrey/olimp/prime"
)

func main() {
	fmt.Println("Raksturīgās kopas")
	kopas.Solve()
	fmt.Println("Divu pirmskaitļu reizinājums")
	prime.Solve()
	fmt.Println("Picērija")
	pizza.Solve()
}
