package main

import (
   "bufio"
    "fmt"
    "os"
    "unicode"
    "regexp"
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

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		a := -1
		b := -1
		for _, c := range line {
			if (unicode.IsDigit(c)) {
				if (a == -1) {
					a = int(c - '0')
				} else {
					b = int(c - '0')
				}
			}
		}
		if (b == -1) {
			b = a
		}
		sum += a*10 + b
	}
	return sum
}


func part2(lines []string) int {
	numMap := map[string]int {
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":  0,
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
	}

	sum := 0
	r := regexp.MustCompile(`^.*?(one|two|three|four|five|six|seven|eight|nine|[0-9]).*$`)
	r2 := regexp.MustCompile(`^.*(one|two|three|four|five|six|seven|eight|nine|[0-9]).*?$`)
	for _, line := range lines {
		v := r.FindStringSubmatch(line)[1]
		v2 := r2.FindStringSubmatch(line)[1]
		a := numMap[v]
		b := numMap[v2]
		sum += a*10 + b
	}
	return sum
}
