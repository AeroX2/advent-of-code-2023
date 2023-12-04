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

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func part1(lines []string) int {
	sum := 0
	for _,line := range lines {
		s := strings.Split(line, ": ")	
		s = strings.Split(s[1], "|")
		winning_numbers_raw := s[0]
		numbers_you_have_raw := s[1]
		numbers_you_have := make(map[int]bool)
		for _,number_raw := range strings.Fields(numbers_you_have_raw) {
			number,err := strconv.Atoi(number_raw)
			check(err)
			numbers_you_have[number] = true
		}

		card_value := -1
		for _,number_raw := range strings.Fields(winning_numbers_raw) {
			number,err := strconv.Atoi(number_raw)
			check(err)
			_, ok := numbers_you_have[number]
			if (ok) {
				card_value += 1
			}
		}

		if (card_value != -1) {
			sum += powInt(2, card_value)
		}

	}
	return sum;
}

func part2(lines []string) int {
	card_amount := make(map[int]int)
	for card_id,line := range lines {
		card_id += 1

		s := strings.Split(line, ": ")	
		s = strings.Split(s[1], "|")
		winning_numbers_raw := s[0]
		numbers_you_have_raw := s[1]
		numbers_you_have := make(map[int]bool)
		for _,number_raw := range strings.Fields(numbers_you_have_raw) {
			number,err := strconv.Atoi(number_raw)
			check(err)
			numbers_you_have[number] = true
		}

		card_value := 0
		for _,number_raw := range strings.Fields(winning_numbers_raw) {
			number,err := strconv.Atoi(number_raw)
			check(err)
			_, ok := numbers_you_have[number]
			if (ok) {
				card_value += 1
			}
		}

		current_amount, ok := card_amount[card_id]
		if (!ok) {
			card_amount[card_id] = 1
			current_amount = 1
		}

		if (card_value != -1) {
			for i := 0; i < card_value; i++ {
				v := card_id+1+i
				_, ok = card_amount[v]
				if (ok) {
					card_amount[v] += current_amount
				} else {
					card_amount[v] = current_amount+1
				}
			}
		}
	}
	sum := 0
	for _,k := range card_amount {
		sum += k
	}
	return sum
}
