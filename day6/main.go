package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// var field = []string{}

type Field struct {
	f        []string
	start    []int
	startDir int
}

func (self *Field) print() {
	fmt.Println("--------------------------------------")
	for _, s := range self.f {
		fmt.Println(s)
	}
	fmt.Println("--------------------------------------")
}

func (self *Field) IsBorderPosition(x, y, dir int) bool {
	return ((x == 0 && dir == dir180) || (y == 0 && dir == dir90) || (x == len(self.f[0])-1 && dir == dir0) || (y == len(self.f)-1 && dir == dir270))
}

func (self *Field) moveForward(x, y, dir int) (int, int, int) {
	xstep, ystep := 0, 0
	switch dir {
	case dir0:
		xstep = 1
	case dir90:
		ystep = -1
	case dir180:
		xstep = -1
	case dir270:
		ystep = 1
	}
	res := 0
	for {
		if y+ystep >= 0 && y+ystep < len(self.f) && x+xstep >= 0 && x+xstep < len(self.f[0]) && self.f[y+ystep][x+xstep] != '#' {
			if self.f[y+ystep][x+xstep] != 'X' {
				self.f[y+ystep] = self.f[y+ystep][:x+xstep] + "X" + self.f[y+ystep][x+xstep+1:]
				res += 1
			}
			x, y = x+xstep, y+ystep
			continue
		}
		break
	}
	return x, y, res
}

const (
	dir0   = 0
	dir90  = 90
	dir180 = 180
	dir270 = 270
)

func turnRight(dir int) int {
	d := []int{dir0, dir270, dir180, dir90}
	for i, v := range d {
		if v != dir {
			continue
		}
		if i < len(d)-1 {
			return d[i+1]
		}
		return 0
	}
	return 0
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	x, y, l := 0, 0, 0
	dir := dir0 // 90, 180, 270
	field := []string{}
	for scanner.Scan() {
		txt := scanner.Text()
		field = append(field, txt)
		expr, _ := regexp.Compile("[<>V^]")
		i := expr.FindIndex([]byte(txt))
		if i != nil {
			x = i[0]
			y = l
			switch txt[i[0]] {
			case '^':
				dir = dir90
			case '<':
				dir = dir180
			case '>':
				dir = dir0
			case 'V':
				dir = dir270

			}
		}
		l += 1
	}
	myField := Field{
		field, []int{x, y}, dir,
	}

	res := 0
	for i := 0; i < len(myField.f); i++ {
		for j := 0; j < len(myField.f[0]); j++ {
			if myField.f[i][j] == '.' {
				myField.f[i] = myField.f[i][:j] + "#" + myField.f[i][j+1:]
				if myField.isloop() {
					res += 1
				}
				myField.f[i] = myField.f[i][:j] + "." + myField.f[i][j+1:]
			}
		}
	}
	fmt.Println(res)
}

func (self *Field) isloop() bool {
	copyF := []string{}
	for i := 0; i < len(self.f); i++ {
		copyF = append(copyF, self.f[i])
	}
	x, y, dir := self.start[0], self.start[1], self.startDir
	copyField := Field{copyF, []int{x, y}, dir}
	copyField.f[y] = copyField.f[y][:x] + "X" + copyField.f[y][x+1:]
	m := map[[3]int]bool{}
	for {
		x, y, _ = copyField.moveForward(x, y, dir)
		if copyField.IsBorderPosition(x, y, dir) {
			return false
		}
		if m[[3]int{x, y, dir}] == true {
			return true
		}
		m[[3]int{x, y, dir}] = true

		dir = turnRight(dir)
	}
}

// func (self *Field) isloop2() bool {
// 	for li, l := range self.f {
// 		for si, s := range l {
// 			if s == '#' {
// 				if self.ispartofloop(si, li) {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// func (self *Field) ispartofloop(x, y int) bool {

// }
