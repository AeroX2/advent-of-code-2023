package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	lines := readFile(os.Args[1])
	fmt.Printf("Part 1 Sum %d\n", part1(lines))
	fmt.Printf("Part 2 Sum %d\n", part2(lines))
}

func readFile(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())
	return lines
}

func lineToNumbers(line string) []int {
	var numbers []int
	for _, possible_num := range strings.Split(line, " ") {
		num, ok := strconv.Atoi(possible_num)
		if ok == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func checkZero(numbers []int) bool {
	for _,n := range numbers {
		if (n != 0) {
			return false
		}
	}
	return true
}

func part1(lines []string) int {
	sum := 0
	for _,line := range lines {
		numbers := lineToNumbers(line)
		var sequences [][]int
		sequences = append(sequences, numbers)
		for !checkZero(sequences[len(sequences)-1]) {
			var new_numbers []int
			for i,n := range numbers[:len(numbers)-1] {
				new_numbers = append(new_numbers, numbers[i+1]-n)
			}
			sequences = append(sequences, new_numbers)
			numbers = new_numbers
		}

		a := sequences[len(sequences)-1]
		b := sequences[len(sequences)-2]
		v := a[len(a)-1] + b[len(b)-1]

		for i:=len(sequences)-3; i>=0; i-- {
			g := sequences[i]
			v += g[len(g)-1]
		}

		sum += v
	}

	return sum
}

func part2(lines []string) int {
	sum := 0
	for _,line := range lines {
		numbers := lineToNumbers(line)
		var sequences [][]int
		sequences = append(sequences, numbers)
		for !checkZero(sequences[len(sequences)-1]) {
			var new_numbers []int
			for i,n := range numbers[:len(numbers)-1] {
				new_numbers = append(new_numbers, numbers[i+1]-n)
			}
			sequences = append(sequences, new_numbers)
			numbers = new_numbers
		}

		a := sequences[len(sequences)-1]
		b := sequences[len(sequences)-2]
		v := b[0] - a[0]

		for i:=len(sequences)-3; i>=0; i-- {
			g := sequences[i]
			v = g[0] - v
		}

		sum += v
	}

	return sum
}
