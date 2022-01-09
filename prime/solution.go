package prime

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	aRawString, _ := reader.ReadString('\n')
	fmt.Print("> ")
	bRawString, _ := reader.ReadString('\n')

	aRawString = strings.Trim(aRawString, "\n")
	bRawString = strings.Trim(bRawString, "\n")

	a, _ := strconv.Atoi(aRawString)
	b, _ := strconv.Atoi(bRawString)
	primes := generatePrimeNumbers(b)

	var answers int
	for i, prime := range primes {
		for _, f := range primes[i:] {
			m := prime * f
			if m >= a && m <= b {
				answers++
			}
		}
	}
	fmt.Printf("Answer: %d\n", answers)
}

func generatePrimeNumbers(max int) []int {
	var primes []int
	min := 2
	max = max / min
	for min <= max {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(min))); i++ {
			if min%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, min)
		}
		min++
	}
	return primes
}
