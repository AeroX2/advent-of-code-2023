package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

// var directions [][]int = [][]int{{0,1}, {1,0}, {0,-1}, {-1,0}, {-1,-1}, {1,1}, {-1,1}, {1,-1}}
var directions [][]int = [][]int{{-1,-1}, {0,-1}, {1,-1}, {-1,0}, {1,0}, {-1,1}, {0,1}, {1,1}}

func part1(lines []string) int {
	sum := 0
	for y,line := range lines {
		num := 0
		valid := false
		for x,c := range line {
			if (unicode.IsDigit(c)) {
				num *= 10
				num += int(c - '0')
				if (!valid) {
					for _,d := range directions {
						nx := x + d[0]
						ny := y + d[1]
						
						if (nx >= 0 && nx < len(lines[0]) && ny >= 0 && ny < len(lines)) {
							r := rune(lines[ny][nx])
							if (r != '.' && !unicode.IsDigit(r)) {
								valid = true
							}
						}
					}
				}
			} else {
				if (valid) {
					sum += num
				}
				num = 0
				valid = false
			}
		}
		if (valid) {
			sum += num
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for y,line := range lines {
		for x,c := range line {
			if (c != '*') {
				continue
			}

			// Would be nice if I had ternery operators in Go.
			var top [][]int
			if (valid(x - 1, y - 1, lines)) {
				top = append(top, []int{x-1, y-1})
			}
			if (valid(x, y - 1, lines)) {
				top = append(top, []int{x, y-1})
			}
			if (valid(x + 1, y - 1, lines)) {
				top = append(top, []int{x+1, y-1})
			}

			var left [][]int
			if (valid(x - 1, y, lines)) {
				left = append(left, []int{x-1,y})
			}

			var right [][]int
			if (valid(x + 1, y, lines)) {
				right = append(right, []int{x+1,y})
			}

			var bottom [][]int
			if (valid(x - 1, y + 1, lines)) {
				bottom = append(bottom, []int{x-1, y+1})
			}
			if (valid(x, y + 1, lines)) {
				bottom = append(bottom, []int{x, y+1})
			}
			if (valid(x + 1, y + 1, lines)) {
				bottom = append(bottom, []int{x+1, y+1})
			}

			valid := 0
			if (len(top) > 0) {
				valid += 1
			}
			if (len(bottom) > 0) {
				valid += 1
			}
			if (len(top) == 2 && top[1][0] - top[0][0] != 1) {
				valid += 1
			}
			if (len(bottom) == 2 && bottom[1][0] - bottom[0][0] != 1) {
				valid += 1
			}
			if (len(left) > 0) {
				valid += 1
			}
			if (len(right) > 0) {
				valid += 1
			}

			if (valid != 2) {
				continue
			}
			var a []int
			if (len(top) > 0) {
				a = top[0]
			} else if (len(bottom) > 0) {
				a = bottom[0]
			} else if (len(left) > 0) {
				a = left[0]
			} else if (len(right) > 0) {
				a = right[0]
			}

			var b []int
			if (len(right) > 0) {
				b = right[0]
			} else if (len(left) > 0) {
				b = left[0]
			} else if (len(bottom) > 0) {
				b = bottom[0]
			} else if (len(top) > 0) {
				b = top[0]
			}

			if (len(top) == 2 && top[1][0] - top[0][0] != 1) {
				a = top[0]
				b = top[1]
			}

			if (len(bottom) == 2 && bottom[1][0] - bottom[0][0] != 1) {
				a = bottom[0]
				b = bottom[1]
			}

			num_1 := expand(a[0], lines[a[1]])
			num_2 := expand(b[0], lines[b[1]])

			sum += num_1 * num_2
		}
	}
	return sum
}

func valid(nx int, ny int, lines []string) bool {
	if (nx >= 0 && nx < len(lines[0]) && ny >= 0 && ny < len(lines)) {
		r := rune(lines[ny][nx])
		return (unicode.IsDigit(r)) 
	}
	return false
}

func expand(x int, line string) int {
	si := x
	for (si >= 0 && unicode.IsDigit(rune(line[si]))) {
		si -= 1
	}

	ei := x
	for (ei < len(line) && unicode.IsDigit(rune(line[ei]))) {
		ei += 1
	}

	v,err := strconv.Atoi(line[si+1:ei])
	check(err)
	return v
}
