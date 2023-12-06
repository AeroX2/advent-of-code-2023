package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
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

func part1(lines []string) int {
	times := lineToNumbers(lines[0])
	distances := lineToNumbers(lines[1])

	sum := 1
	for i,time_i := range times {
		time := float64(time_i)
		distance := float64(distances[i])
		lb := math.Floor(0.5*(time-math.Sqrt(math.Pow(time,2)-4*distance)))
		ub := math.Ceil(0.5*((math.Sqrt(math.Pow(time,2)-4*distance)+time)))
		sum *= int(ub) - int(lb) - 1
	}

	return sum
}

func part2(lines []string) int {
	times := lineToNumbers(
		strings.Replace(
			strings.Split(lines[0], ":")[1],
			" ",
			"",
			-1,
		),
	)

	distances := lineToNumbers(
		strings.Replace(
			strings.Split(lines[1], ":")[1],
			" ",
			"",
			-1,
		),
	)

	sum := 1
	for i,time_i := range times {
		time := float64(time_i)
		distance := float64(distances[i])
		lb := math.Floor(0.5*(time-math.Sqrt(math.Pow(time,2)-4*distance)))
		ub := math.Ceil(0.5*((math.Sqrt(math.Pow(time,2)-4*distance)+time)))
		sum *= int(ub) - int(lb) - 1
	}

	return sum
}
