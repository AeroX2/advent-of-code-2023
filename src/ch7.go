package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
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

type CardHand struct {
	hand string
	bet int
	_rank int
}

func (ch CardHand) rank(enable_joker bool) int {
	m := make(map[rune]int)
	for _,c := range ch.hand {
		_, ok := m[c]
		if (ok) {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	var v []int
	for c,x := range m {
		if (c != 'J') {
			v = append(v, x)
		}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})

	jokers := 0
	j, ok := m['J']
	if (ok && enable_joker) {
		jokers = j
	}

	if (len(v) == 0 || v[0]+jokers == 5) {
		return 6
	} else if (v[0]+jokers == 4) {
		return 5
	} else if (v[0]+jokers == 3 && len(v) > 1 && v[1] == 2) {
		return 4
	} else if (v[0]+jokers == 3) {
		return 3
	} else if (v[0]+jokers == 2 && len(v) > 1 && v[1] == 2) {
		return 2
	} else if (v[0]+jokers == 2) {
		return 1
	} else {
		return 0
	}
}

var strength string = "AKQT98765432J"

func calc(lines []string, enable_joker bool) int {
	var card_hands []CardHand
	for _,line := range lines {
		split := strings.Split(line, " ")
		hand := split[0]
		bet, err := strconv.Atoi(split[1])
		check(err)

		card_hands = append(card_hands, CardHand{hand, bet, -1})
	}

	sort.Slice(card_hands, func(i, j int) bool {
		a := card_hands[i]
		b := card_hands[j]
		ar := a.rank(enable_joker)
		br := b.rank(enable_joker)

		if (ar != br) {
			return ar < br
		}

		for i,c := range a.hand {
			aa := strings.Index(strength, string(c))
			bb := strings.Index(strength, string(b.hand[i]))
			if (aa != bb) {
				return aa > bb
			}
		}
		return true
	})

	sum := 0
	for i,ch := range card_hands {
		sum += (i+1) * ch.bet
	}

	return sum
}

func part1(lines []string) int {
	return calc(lines, false)
}

func part2(lines []string) int {
	return calc(lines, true)
}

