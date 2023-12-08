package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Node struct {
	left string
	right string
}

func findEnd(instructions string, graph map[string]Node, current_element string) int {
	step := 0
	for string(current_element[len(current_element)-1]) != "Z" {
		instruction := string(instructions[step % len(instructions)])
		if (instruction == "L") {
			current_element = graph[current_element].left
		} else {
			current_element = graph[current_element].right
		}
		step += 1
	}
	return step
}

func part1(lines []string) int {
	instructions := lines[0]

	graph := make(map[string]Node)
	for _,line := range lines[2:] {
		x := strings.Split(line, " = ")
		p := x[0]
		y := strings.Trim(x[1], "()")
		z := strings.Split(y, ", ")
		a := z[0]
		b := z[1]

		graph[p] = Node{a,b}
	}

	return findEnd(instructions, graph, "AAA")
}

func checkEnd(l []string) bool {
	for _,n := range l {
		if (string(n[len(n)-1]) != "Z") {
			return false
		}
	} 
	return true
}

// gcd computes the greatest common divisor of two integers using Euclid's algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm computes the least common multiple of two integers using the formula: lcm(a, b) = |a * b| / gcd(a, b)
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	g := gcd(a, b)
	return (a / g) * b
}

// lcmList computes the least common multiple of a list of integers.
func lcmList(numbers []int) int {
	result := 1

	for _, num := range numbers {
		result = lcm(result, num)
	}

	return result
}

func part2(lines []string) int {
	instructions := lines[0]

	graph := make(map[string]Node)
	for _,line := range lines[2:] {
		x := strings.Split(line, " = ")
		p := x[0]
		y := strings.Trim(x[1], "()")
		z := strings.Split(y, ", ")
		a := z[0]
		b := z[1]

		graph[p] = Node{a,b}
	}

	var starts []string
	for k,_ := range graph {
		if (string(k[len(k)-1]) == "A") {
			starts = append(starts, k)
		}
	}

	var cycles []int
	for _,start := range starts {
		cycles = append(cycles, findEnd(instructions, graph, start))
	}

	return lcmList(cycles)
}
