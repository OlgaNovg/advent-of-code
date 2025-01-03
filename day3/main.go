package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "regexp"
)

func mul(s string) int {
	rdig, _ := regexp.Compile("[0-9]+")
	numbers := rdig.FindAllString(s, 2)
	n1, _ := strconv.Atoi(numbers[0])
	n2, _ := strconv.Atoi(numbers[1])
	return n1 * n2
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	res := 0
	do := true
	for scanner.Scan() {
		l := scanner.Text()
		dostr, dontstr := "do[(][])]", "don't[(][])]"
		rWhole, _ := regexp.Compile("(mul[(][0-9]+,[0-9]+[)])|((" + dostr + ")|(" + dontstr + "))")
		doreg, _ := regexp.Compile(dostr)
		dontreg, _ := regexp.Compile(dontstr)
		for _, s := range rWhole.FindAllString(l, -1) {
			// fmt.Println(s)
			if doreg.MatchString(s) {
				do = true
				continue
			}
			if dontreg.MatchString(s) {
				do = false
				continue
			}
			if do {
				res = res + mul(s)
			}
		}
	}
	fmt.Println(res)
}
