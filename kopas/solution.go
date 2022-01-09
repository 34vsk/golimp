package kopas

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	a, b := parseUserInput()
	for ai := 1; ai < len(a); ai++ {
		for bi := 1; bi < len(b); bi++ {
			if equal(sort(removeDuplicates(a[:ai])), sort(removeDuplicates(b[:bi]))) {
				fmt.Printf("Answer: %d %d\n", ai, bi)
				return
			}
		}
	}
	fmt.Printf("Answer: 0 0\n")
}

func removeDuplicates(source []int) []int {
	k := make(map[int]bool)
	t := []int{}
	for _, entry := range source {
		if _, exists := k[entry]; !exists {
			k[entry] = true
			t = append(t, entry)
		}
	}
	return t
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func sort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func parseUserInput() ([]int, []int) {
	var a, b []int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	aRawString, _ := reader.ReadString('\n')
	fmt.Print("> ")
	bRawString, _ := reader.ReadString('\n')

	aRawString = strings.Trim(aRawString, "\n")
	bRawString = strings.Trim(bRawString, "\n")

	aRawSlice := strings.Split(aRawString, " ")
	bRawSlice := strings.Split(bRawString, " ")

	for _, el := range aRawSlice {
		d, err := strconv.Atoi(el)
		if err != nil {
			panic("malformed user input (failed to parse digit)")
		}
		a = append(a, d)
	}

	for _, el := range bRawSlice {
		d, err := strconv.Atoi(el)
		if err != nil {
			panic("malformed user input (failed to parse digit)")
		}
		b = append(b, d)
	}
	return a[1:], b[1:]
}
