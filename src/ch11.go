package main

import (
	"bufio"
	"fmt"
	"os"
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

type Coord struct {
	x     int
	y     int
	num int
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func intMin(a, b int) int {
	if (a < b){
		return a
	}
	return b
}

func intMax(a, b int) int {
	if (a > b){
		return a
	}
	return b
}

func calc(lines []string, expansion int) int {
	g := 1
	var galaxies []Coord
	galaxy_xs := make(map[int]bool)
	galaxy_ys := make(map[int]bool)
	for y,line := range lines {
		for x,c := range line {
			if (c == '#') {
				galaxies = append(galaxies, Coord{x,y,g})
				g += 1
				galaxy_xs[x] = true
				galaxy_ys[y] = true
			}
		}
	}

	sum := 0
	for i,galaxy1 := range galaxies {
		// min_distance := -1
		for _,galaxy2 := range galaxies[i+1:] {
			if (galaxy1 == galaxy2) {
				continue
			}

			distance := intAbs(galaxy1.x-galaxy2.x) + intAbs(galaxy1.y-galaxy2.y)
			for x:=intMin(galaxy1.x, galaxy2.x); x<intMax(galaxy1.x, galaxy2.x); x++ {
				_, ok := galaxy_xs[x]
				if (!ok) {
					distance += expansion-1
				}
			}
			for y:=intMin(galaxy1.y, galaxy2.y); y<intMax(galaxy1.y, galaxy2.y); y++ {
				_, ok := galaxy_ys[y]
				if (!ok) {
					distance += expansion-1
				}
			}

			// if (min_distance == -1 || distance < min_distance) {
			// 	fmt.Println(galaxy1, galaxy2)
			// 	min_distance = distance
			// }
			sum += distance
		}
		// fmt.Println(min_distance)
		// sum += min_distance
	}
	return sum
}

func part1(lines []string) int {
	return calc(lines, 2)
}

func part2(lines []string) int {
	return calc(lines, 100000)
}
