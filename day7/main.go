package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type data struct {
	res  int
	nums []int
}

var (
	plus = "+"
	mult = "*"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	inputset := []data{}
	for scanner.Scan() {
		txt := scanner.Text()
		res := strings.Split(txt, ":")
		d := data{}
		d.res, _ = strconv.Atoi(res[0])

		for _, s := range strings.Split(res[1], " ") {
			i, _ := strconv.Atoi(s)
			d.nums = append(d.nums, i)
		}
		inputset = append(inputset, d)
	}
	ret := 0
	for _, i := range inputset {
		varres := i.calculate(len(i.nums) - 1)
		if slices.Contains(varres, i.res) {
			ret += i.res
		} else {
			fmt.Println(i)
		}

	}
	fmt.Println(ret)
}

func (self *data) calculate(i int) []int {
	if i == 0 {
		return []int{self.nums[0]}
	}
	res := self.calculate(i - 1)
	ret := []int{}
	for _, r := range res {
		if self.nums[i]+r <= self.res {
			ret = append(ret, self.nums[i]+r)
		}
		if self.nums[i]*r <= self.res {
			ret = append(ret, self.nums[i]*r)
		}
	}
	return ret
}
