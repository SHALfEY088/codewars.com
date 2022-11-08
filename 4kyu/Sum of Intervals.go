// https://www.codewars.com/kata/52b7ed099cdc285c300001cd/train/go
package main

import (
	"fmt"
	"sort"
)

const (
	l = 0
	r = 1
)

func main() { //       0       1        2       3		   4
	input := [][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}}
	SumOfIntervals(input)
}

func SumOfIntervals(intervals [][2]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	length := intervals[0][r] - intervals[0][l]

	point := intervals[0][r]

	//fmt.Println(len(intervals))

	for i := 1; i < len(intervals); i++ {
		if point >= intervals[i][r] {
			continue
		} else if point >= intervals[i][l] {
			length += intervals[i][r] - point
			point = intervals[i][r]
		} else {
			length += intervals[i][r] - intervals[i][l]
			point = intervals[i][r]
		}
	}

	// fmt.Println(intervals)
	// fmt.Println(length)

	return length
}

func SumOfIntervals2(intervals [][2]int) int {
	fmt.Println(intervals)
	added := make([][2]int, 0)
	i := 0
	for {
		start, end := intervals[i][0], intervals[i][1]
		g, ok, in := check(added, start, end)
		if ok {
			f(&added[g], start, end)
		} else if in {
			goto next
		} else {
			added = append(added, [2]int{start, end})
		}
	next:
		i++
		if i == len(intervals) {
			if len(added) == len(intervals) {
				break
			} else {
				i = 0
				intervals = added
				added = make([][2]int, 0)
			}
		}
	}
	var sum int
	for i := 0; i < len(added); i++ {
		sum += added[i][1] - added[i][0]
	}
	return sum
}

func f(ai *[2]int, ss, es int) {
	sf, ef := ai[0], ai[1]
	switch {
	// sfss-------ss-------ssef------------------es
	case ss <= ef && ss >= sf && es > ef:
		ai[1] = es
	// sssf------------sf-----------sfes-------------ef
	case ss <= sf && sf <= es && ef >= es:
		ai[0] = ss
	// ss-----------sf-----------------ef-------------es
	case ss <= sf && ef <= es:
		ai[0], ai[1] = ss, es
	}
}

func check(added [][2]int, start, end int) (int, bool, bool) {
	for i := 0; i < len(added); i++ {
		if added[i][0] >= start && added[i][0] <= end ||
			added[i][1] >= start && added[i][1] <= end {
			return i, true, false
		}
		if start >= added[i][0] && start <= added[i][1] &&
			end >= added[i][0] && end <= added[i][1] {
			return 0, false, true
		}
	}
	return 0, false, false
}

// func SumOfIntervals(intervals [][2]int) int {

// 	mp := make(map[int]int)

// 	for _, v := range intervals {
// 		for i := v[0]; i < v[1]; i++ {
// 			mp[i]++
// 		}
// 	}

// 	return len(mp)
// }
