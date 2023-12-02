package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

var a = make(map[string]int)

type fn1 func (diceValue int, diceName string) bool
type fn2 func (gameId int, valid bool)
type fn3 func ()

func parse(lines []string, fn1 fn1, fn2 fn2, fn3 fn3) {
	r := regexp.MustCompile(`Game (\d+): (.+)`)
	r2 := regexp.MustCompile(`([^;]+;?)`)
	r3 := regexp.MustCompile(` ?(\d+) (\w+)(,|;)?`)

	for _,line := range lines  {
		fn3()

		m := r.FindStringSubmatch(line)
		gameId, err := strconv.Atoi(m[1])
		check(err)

		dices := m[2]

		valid := true
		out:
		for _, diceSet := range r2.FindAllStringSubmatch(dices, -1) {
			for _, dice := range r3.FindAllStringSubmatch(diceSet[1], -1) {
				diceValue, err := strconv.Atoi(dice[1])
				check(err)
				diceName := dice[2]

				valid = fn1(diceValue, diceName)
				if (!valid) {
					break out
				}
			}
		}

		fn2(gameId, valid)
	}
}

func part1(lines []string) int {
	sum := 0
	parse(lines, func (diceValue int, diceName string) bool {
		v := -1
		switch diceName {
			case "red":
				v = 12
			case "green":
				v = 13
			case "blue":
				v = 14
		}
		return diceValue <= v
	}, func (gameId int, valid bool) {
		if (valid) {
			sum += gameId
		}
	}, func() { 
	})
	return sum
}

func part2(lines []string) int {
	diceMap := make(map[string]int)
	sum := 0

	parse(lines, func (diceValue int, diceName string) bool {
		if (diceValue > diceMap[diceName]) {
			diceMap[diceName] = diceValue
		}
		return true
	}, func (gameId int, valid bool) {
		sum += diceMap["red"] * diceMap["green"] * diceMap["blue"]
	}, func () {
		diceMap["red"] = 0
		diceMap["green"] = 0
		diceMap["blue"] = 0
	})

	return sum
}
