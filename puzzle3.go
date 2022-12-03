package main

import (
	"bufio"
	"fmt"
	"os"
)

func findRepeated(left string, right string) string {
	for _, l1 := range left {
		for _, l2 := range right {
			if l1 == l2 {
				return string(l1)
			}
		}
	}

	return ""
}

func calc(input string) int {
	left, right := input[0:len(input)/2], input[len(input)/2:]
	repeated := findRepeated(left, right)

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

	for scanner.Scan() {
		value += calc(scanner.Text())
	}

	fmt.Println(value)
}
