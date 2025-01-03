package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// type rule []int

var rules = [][]int{}

// type update []int

var pagerule = make(map[int][]int)
var updates = [][]int{}
var pagepos = []map[int]int{}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		parse(scanner.Text())
	}
	unordered := []int{}
	res := 0
	for i, u := range updates {
		for _, p := range u {
			if !checkrules(i, p) {
				unordered = append(unordered, i)
				goto cont
			}
		}
		res += u[len(u)/2]
	cont:
		continue
	}
	fmt.Println(res)

	res = 0
	for _, ui := range unordered {
		u := updates[ui]
		fmt.Println(u)
		slices.SortFunc(u, isordered)
		fmt.Println(u)
		res += u[len(u)/2]
	}
	fmt.Println(res)

}

func isordered(p1, p2 int) int {
	pagerules := pagerule[p1]
	for _, ri := range pagerules {
		r := rules[ri]
		n1, n2 := r[0], r[1]
		if p1 == 13 && p2 == 29 {
		}
		if !((n1 == p1 && n2 == p2) || (n1 == p2 && n2 == p1)) {
			continue
		}
		if n1 == p1 && n2 == p2 {
			return 1
		}
		return -1
	}
	return 0
}

func addruletopage(p, r int) {
	if _, ok := pagerule[p]; !ok {
		pagerule[p] = []int{}
	}
	pr := pagerule[p]
	pr = append(pr, r)
	pagerule[p] = pr
}

func parse(s string) {
	if s == "" {
		return
	}

	if ok, _ := regexp.MatchString("[0-9]+[|][0-9]+", s); ok {
		rule := strings.Split(s, "|")
		n1, _ := strconv.Atoi(rule[0])
		n2, _ := strconv.Atoi(rule[1])
		rules = append(rules, []int{n1, n2})
		addruletopage(n1, len(rules)-1)
		addruletopage(n2, len(rules)-1)
		return
	}
	pages := strings.Split(s, ",")
	update := []int{}
	pagepos = append(pagepos, map[int]int{})
	for i, p := range pages {
		pint, _ := strconv.Atoi(p)
		update = append(update, pint)
		pagepos[len(pagepos)-1][pint] = i
	}
	updates = append(updates, update)
}

func checkrules(i, p int) bool {
	pagerules := pagerule[p]
	ruleschecked := make([]bool, len(pagerules))

	for j, ri := range pagerules {

		if ruleschecked[j] {
			continue
		}
		rule := rules[ri]
		n1pos, ok1 := pagepos[i][rule[0]]
		n2pos, ok2 := pagepos[i][rule[1]]

		if !(ok1 && ok2) {
			ruleschecked[j] = true
		} else if n1pos > n2pos {
			return false
		} else {
			ruleschecked[j] = true
		}
	}
	return true
}
