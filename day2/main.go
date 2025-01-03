package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func IsSafe(line []string) bool {
	t := 0
	for i := 1; i < len(line); i++ {
		n, _ := strconv.Atoi(line[i])
		n0, _ := strconv.Atoi(line[i-1])
		if i == 1 {
			if n > n0 {
				t = 1
			} else if n < n0 {
				t = -1
			} else {
				return false
			}
		} else {
			if !((n > n0 && t == 1) || (n < n0 && t == -1)) {
				return false
			}
		}
		if math.Abs(float64(n)-float64(n0)) <= 0 || math.Abs(float64(n)-float64(n0)) > 3 {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if IsSafe(line) {
			res = res + 1
		} else {
			for i := 0; i < len(line); i++ {
				a := []string{}
				a = append(a, line[:i]...)
				a = append(a, line[i+1:]...)
				if IsSafe(a) {
					res = res + 1
					break
				}
			}
		}
	}
	fmt.Println(res)
}
