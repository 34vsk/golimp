package pizza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	rawTotalParks, _ := reader.ReadString('\n')
	rawTotalParks = strings.Trim(rawTotalParks, "\n")
	totalParks := parseInt(rawTotalParks)

	parks := make([]int, totalParks)

	for i := 1; i < totalParks; i++ {
		fmt.Print("> ")
		rawInput, _ := reader.ReadString('\n')
		rawInput = strings.Trim(rawInput, "\n")
		s := strings.Split(rawInput, " ")
		a, b := parseInt(s[0]), parseInt(s[1])
		parks[a-1]++
		parks[b-1]++
	}

	var max, targetIndex int
	for i, v := range parks {
		if v > max {
			max = v
			targetIndex = i
		}
	}
	fmt.Printf("Answer: %d\n", targetIndex+1)
}

func parseInt(source string) int {
	t, _ := strconv.Atoi(source)
	return t
}
