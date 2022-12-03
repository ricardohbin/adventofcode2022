package main

import (
	"bufio"
	"fmt"
	"os"
)

func findRepeated(first string, second string, third string) string {
	for _, l1 := range first {
		for _, l2 := range second {
			if l1 == l2 {
				for _, l3 := range third {
					if l1 == l3 {
						return string(l1)
					}
				}
			}
		}
	}

	return ""
}

func Calc(input []string) int {
	repeated := findRepeated(input[0], input[1], input[2])

	charcode := int(repeated[0])
	var priority int

	// using charcode to help with the priorities:
	// 'A' char is 65 charcode and 27 priority, so the priority to uppercase can be: "charcode - 38 == priority"
	// while `a` is 97 and need to be 1 priority, so the priority to lowercase can be "charcode - 96 == priority"
	if charcode < 97 { //uppercase
		priority = charcode - 38
	} else { //lowercase
		priority = charcode - 96
	}

	return priority
}

func main() {
	f, err := os.Open("3.input")

	if err != nil {
		panic("can't open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	value := 0

	var m []string

	for scanner.Scan() {
		m = append(m, scanner.Text())

		if len(m) == 3 {
			value += Calc(m)
			m = nil
		}
	}

	fmt.Println(value)
}
