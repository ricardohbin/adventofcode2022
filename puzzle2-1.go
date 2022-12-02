package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(condition string) int {
	table := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	return table[condition]
}

func change(opponent string, value int) string {
	table := map[string]map[int]string{
		"A": {
			3: "X",
			6: "Y",
			0: "Z",
		},
		"B": {
			0: "X",
			3: "Y",
			6: "Z",
		},
		"C": {
			6: "X",
			0: "Y",
			3: "Z",
		},
	}
	return table[opponent][value]
}

func main() {
	f, err := os.Open("2.input")

	if err != nil {
		panic("can't open file")
	}

	defer f.Close()

	bonusWin := map[string]int{"X": 1, "Y": 2, "Z": 3}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		value := scanner.Text()
		opponent := string(value[0])
		condition := string(value[2])
		checkValue := check(condition)
		total += checkValue + bonusWin[change(opponent, checkValue)]
	}

	fmt.Println(total)
}
