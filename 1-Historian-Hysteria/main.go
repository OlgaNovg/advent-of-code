package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MinHeap struct {
	tree []int
}

func Left(index int) int {
	index += 1
	return 2*index - 1
}

func Right(index int) int {
	index += 1
	return 2 * index
}

func Parent(index int) int {
	index += 1
	return index/2 - 1
}

func (h *MinHeap) Insert(v int) {
	h.tree = append(h.tree, v)
	if len(h.tree) == 1 {
		return
	}
	i := len(h.tree) - 1
	p := Parent(i)
	for i > 0 && p >= 0 && h.tree[p] > h.tree[i] {
		h.tree[p], h.tree[i] = h.tree[i], h.tree[p]
		i = p
		p = Parent(i)
	}
}
func (h *MinHeap) MinHeapifyRoot() {
	for i, l, r := 0, Left(0), Right(0); (l < len(h.tree) && h.tree[i] > h.tree[l]) || (r < len(h.tree) && h.tree[i] > h.tree[r]); l, r = Left(i), Right(i) {
		if r >= len(h.tree) || (l < len(h.tree) && h.tree[r] > h.tree[l]) {
			h.tree[i], h.tree[l] = h.tree[l], h.tree[i]
			i = l
		} else {
			h.tree[i], h.tree[r] = h.tree[r], h.tree[i]
			i = r
		}
	}
}

func (h *MinHeap) GetMin() (v int) {
	if len(h.tree) == 0 {
		return 0
	}
	v = h.tree[0]
	h.tree[0] = h.tree[len(h.tree)-1]
	h.tree = h.tree[:len(h.tree)-1]
	h.MinHeapifyRoot()
	return
}

func part1() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	res := 0
	var myheap1 MinHeap
	var myheap2 MinHeap
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		n1, _ := strconv.Atoi(line[0])
		n2, _ := strconv.Atoi(line[1])
		myheap1.Insert(n1)
		myheap2.Insert(n2)
	}

	for len(myheap1.tree) > 0 {
		n1, n2 := myheap1.GetMin(), myheap2.GetMin()
		if n1 > n2 {
			res = res + n1 - n2
		} else {
			res = res + n2 - n1
		}
	}

	fmt.Println(res)
}

func part2() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	arr1, arr2 := []int{}, []int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		n1, _ := strconv.Atoi(line[0])
		n2, _ := strconv.Atoi(line[1])
		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}
	m := make(map[int][]int)
	for _, x := range arr1 {
		if v, ok := m[x]; ok {
			v[0] = v[0] + 1
		} else {
			m[x] = []int{1, 0}
		}
	}

	for _, x := range arr2 {
		if v, ok := m[x]; ok {
			v[1] = v[1] + 1
		}
	}

	res := 0
	for k, v := range m {
		res = res + k*v[0]*v[1]
	}
	fmt.Println(res)

}

func main() {
	part2()
}
