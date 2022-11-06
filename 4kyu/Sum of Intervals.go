// https://www.codewars.com/kata/52b7ed099cdc285c300001cd/train/go
package main

import (
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

	sort.SliceStable(intervals, func(i, j int) bool {
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

// func SumOfIntervals(intervals [][2]int) int {

// 	mp := make(map[int]int)

// 	for _, v := range intervals {
// 		for i := v[0]; i < v[1]; i++ {
// 			mp[i]++
// 		}
// 	}

// 	return len(mp)
// }
