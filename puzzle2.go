package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(opponent string, mine string) int {
	table := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	return table[opponent][mine]
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
		my := string(value[2])
		total += check(opponent, my) + bonusWin[my]
	}

	fmt.Println(total)
}
