package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("1.input")

	if err != nil {
		panic("can't open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	counter := 0

	var values []int

	for scanner.Scan() {
		if scanner.Text() == "" {
			values = append(values, counter)
			counter = 0
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic("Invalid text. error: " + err.Error())
			} else {
				counter += value
			}
		}
	}

	sort.Ints(values)

	last := len(values) - 1
	lastValues := 0

	for _, v := range values[last-2:] {
		lastValues += v
	}

	fmt.Println(lastValues)
}
