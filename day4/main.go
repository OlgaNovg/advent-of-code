package main

import (
	"bufio"
	"fmt"
	"os"
)

type matrix []string

var moves = make(map[string][2]int)

func (m *matrix) checkxmas(i, j int) int {
	res := 0
	for k := range moves {
		res += m.checkdirection(i, j, k, "XMAS")
	}
	return res
}

func (m matrix) checkxmas2(i, j int) int {
	res := 0
	c1 := [2][2]int{{-1, -1}, {1, 1}}
	c2 := [2][2]int{{1, -1}, {-1, 1}}

	counter := 0
	c := c1
	for {
		counter += 1
		if counter > 2 {
			break
		}
		i1, j1 := i+c[0][0], j+c[0][1]
		i2, j2 := i+c[1][0], j+c[1][1]
		if i1 < 0 || i1 >= len(m[0]) || i2 < 0 || i2 >= len(m[0]) || j1 < 0 || j1 >= len(m) || j2 < 0 || j2 >= len(m) {
			continue
		}
		l1, l2 := m[i1][j1], m[i2][j2]
		if (l1 == 'M' && l2 == 'S') || (l1 == 'S' && l2 == 'M') {
			res += 1
		}
		c = c2
	}
	if res == 2 {
		return 1
	}
	return 0
}

func (m matrix) checkdirection(i, j int, direction string, word string) int {
	if i < 0 || i >= len(m[0]) || j < 0 || j >= len(m) {
		return 0
	}
	if word == string(m[i][j]) {
		return 1
	}
	if word[0] == m[i][j] {
		return m.checkdirection(i+moves[direction][0], j+moves[direction][1], direction, word[1:])
	}
	return 0

}

func main() {
	moves["l"] = [2]int{-1, 0}
	moves["r"] = [2]int{1, 0}
	moves["u"] = [2]int{0, -1}
	moves["d"] = [2]int{0, 1}
	moves["lu"] = [2]int{-1, -1}
	moves["ru"] = [2]int{1, -1}
	moves["ld"] = [2]int{-1, 1}
	moves["rd"] = [2]int{1, 1}

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var m matrix
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}

	res := 0
	for i, l := range m {
		for j, s := range l {
			if s == 'X' {
				res += m.checkxmas(i, j)
			}
		}
	}
	fmt.Println(res)

	res = 0
	for i, l := range m {
		for j, s := range l {
			if s == 'A' {
				res += m.checkxmas2(i, j)
			}
		}
	}
	fmt.Println(res)
}
