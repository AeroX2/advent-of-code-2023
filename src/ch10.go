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
	depth int
}

func findStart(lines []string) Coord {
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				return Coord{x, y, 0}
			}
		}
	}
	return Coord{-1, -1, 0}
}

func getDirections(lines []string, current Coord) []Coord {
	var directions []Coord

	switch lines[current.y][current.x] {
	case '|':
		directions = append(directions, Coord{0, -1, 0})
		directions = append(directions, Coord{0, 1, 0})
	case '-':
		directions = append(directions, Coord{-1, 0, 0})
		directions = append(directions, Coord{1, 0, 0})
	case 'L':
		directions = append(directions, Coord{0, -1, 0})
		directions = append(directions, Coord{1, 0, 0})
	case 'J':
		directions = append(directions, Coord{0, -1, 0})
		directions = append(directions, Coord{-1, 0, 0})
	case '7':
		directions = append(directions, Coord{0, 1, 0})
		directions = append(directions, Coord{-1, 0, 0})
	case 'F':
		directions = append(directions, Coord{0, 1, 0})
		directions = append(directions, Coord{1, 0, 0})
	}

	if lines[current.y][current.x] == 'S' {
		v := lines[current.y-1][current.x]
		if v == '|' || v == '7' || v == 'F' {
			return []Coord{{0,-1,0}}
		}
		v = lines[current.y+1][current.x]
		if v == '|' || v == 'L' || v == 'J' {
			return []Coord{{0,1,0}}
		}
		v = lines[current.y][current.x-1]
		if v == '-' || v == 'F' || v == 'L' {
			return []Coord{{-1,0,0}}
		}
		v = lines[current.y][current.x+1]
		if v == '-' || v == '7' || v == 'J' {
			return []Coord{{1,0,0}}
		}
		panic("This shouldn't ever happen")
	}

	return directions
}

var seen map[Coord]bool
var loop []Coord

func part1(lines []string) int {
	max_depth := 0

	start := findStart(lines)

	seen = make(map[Coord]bool)
	seen[start] = true
	loop = append(loop, start)

	var stack []Coord
	stack = append(stack, start)
	for len(stack) > 0 {

		var current Coord
		current, stack = stack[0], stack[1:]
		if current.depth > max_depth {
			max_depth = current.depth
		}

		directions := getDirections(lines, current)

		for _, direction := range directions {
			x := current.x + direction.x
			y := current.y + direction.y
			if x < 0 || x > len(lines[0]) ||
				y < 0 || y > len(lines) {
				continue
			}

			_, ok := seen[Coord{x, y, 0}]
			if ok {
				continue
			}
			seen[Coord{x, y, 0}] = true

			if lines[y][x] != '.' {
				stack = append(stack, Coord{x, y, current.depth + 1})
				loop = append(loop, Coord{x,y,0})
			}
		}

	}

	// for y := 0; y < len(lines); y++ {
	// 	for x := 0; x < len(lines[0]); x++ {
	// 		// _, ok := seen[Coord{x, y, 0}]
	// 		var found Coord
	// 		g := 0
	// 		for i,l := range loop {
	// 			if (l.x == x && l.y == y) {
	// 				found = Coord{x,y,0}
	// 				g = i
	// 				break
	// 			}
	// 		}

	// 		if (found != Coord{}) {
	// 			fmt.Printf("%-10d", g)
	// 		} else {
	// 			fmt.Print(".         ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return max_depth
}

func windingNumberTest(point Coord, polygon []Coord) int {
	x, y := point.x, point.y
	windingNumber := 0

	for i := 0; i < len(polygon); i++ {
		x1, y1 := polygon[i].x, polygon[i].y
		x2, y2 := polygon[(i+1)%len(polygon)].x, polygon[(i+1)%len(polygon)].y

		// Check if the point is below or above the edge (adjusted for y-axis pointing downwards)
		if y1 >= y {
			if y2 < y && isLeft(x1, y1, x2, y2, x, y) > 0 {
				windingNumber++
			}
		} else if y2 >= y && isLeft(x1, y1, x2, y2, x, y) < 0 {
			windingNumber--
		}
	}

	return windingNumber
}

func isLeft(x1, y1, x2, y2, x, y int) int {
	return (x2-x1)*(y-y1) - (x-x1)*(y2-y1)
}

func part2(lines []string) int {
	var spaces []Coord
	for y,line := range lines {
		for x,_ := range line {
			_, ok := seen[Coord{x,y,0}]
			if (!ok) {
				spaces = append(spaces, Coord{x,y,0})
			}
		}
	}

	sum := 0
	for _,space := range spaces {
		if windingNumberTest(space, loop) != 0 {
			sum += 1
		}
	}
	
	return sum
}
