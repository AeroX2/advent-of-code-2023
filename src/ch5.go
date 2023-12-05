package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data_raw, err := os.ReadFile(os.Args[1])
	check(err)
	data := string(data_raw)

	fmt.Printf("Part 1 Lowest %d\n", part1(data))
	fmt.Printf("Part 2 Lowest %d\n", part2(data))
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

type Range struct {
	src_min  int
	src_max  int
	dest_min int
	dest_max int
}

func (r Range) convert(num int) (int, bool) {
	if num >= r.src_min && num <= r.src_max {
		v := (num - r.src_min) + r.dest_min
		return v, true
	}
	return 0, false
}

func toRanges(ranges_str string) []Range {
	var ranges []Range
	for _, v := range strings.Split(ranges_str, "\n")[1:] {
		n := lineToNumbers(v)
		r := Range{
			src_min:  n[1],
			src_max:  n[1] + n[2],
			dest_min: n[0],
			dest_max: n[0] + n[2],
		}
		ranges = append(ranges, r)
	}
	return ranges
}

func part1(data string) int {
	maps := strings.Split(data, "\n\n")
	seeds := lineToNumbers(maps[0])

	var ranges_maps [][]Range
	for _, ranges_str := range maps[1:] {
		ranges_maps = append(ranges_maps, toRanges(ranges_str))
	}

	lowest_seed := -1
	for _, seed := range seeds {
		for _, ranges := range ranges_maps {
			for _, r := range ranges {
				new_seed, ok := r.convert(seed)
				if ok {
					seed = new_seed
					break
				}
			}
		}

		if lowest_seed == -1 || seed < lowest_seed {
			lowest_seed = seed
		}
	}

	return lowest_seed
}

func part2(data string) int {
	maps := strings.Split(data, "\n\n")
	seeds_v := lineToNumbers(maps[0])

	var ranges_maps [][]Range
	for _, ranges_str := range maps[1:] {
		ranges_maps = append(ranges_maps, toRanges(ranges_str))
	}

	lowest_seed := -1
	for i := 0; i < len(seeds_v); i += 2 {
		n := seeds_v[i : i+2]
		////seed_ranges = append(seed_ranges, Range {
		//	src_min: n[0],
		//	src_max: n[0]+n[1],
		//	dest_min: n[0],
		//	dest_max: n[0]+n[1],
		//})

		for seed := n[0]; seed < n[0]+n[1]; seed++ {
			a := seed
			for _, ranges := range ranges_maps {
				for _, r := range ranges {
					new_seed, ok := r.convert(a)
					if ok {
						a = new_seed
						break
					}
				}
			}

			if (a != -1) {
				if lowest_seed == -1 || a < lowest_seed {
					lowest_seed = a
				}
			}
		}
	}

	return lowest_seed
}
