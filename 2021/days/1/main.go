package main

import (
	"fmt"

	"github.com/dewski/adventofcode/2021/inputs"
)

// 199 (N/A - no previous measurement)
// 200 (increased)
// 208 (increased)
// 210 (increased)
// 200 (decreased)
// 207 (increased)
// 240 (increased)
// 269 (increased)
// 260 (decreased)
// 263 (increased)
func increased(nums []int) int {
	// no nums provided, return
	if len(nums) == 0 {
		return 0
	}

	// Need to ignore initial depth when counting number of times depth increases
	previousNum := nums[0]
	increasedCounter := 0

	for _, num := range nums {
		if num > previousNum {
			increasedCounter += 1
		}

		previousNum = num
	}

	return increasedCounter
}

// 199  A
// 200  A B
// 208  A B C
// 210    B C D
// 200  E   C D
// 207  E F   D
// 240  E F G
// 269    F G H
// 260      G H
// 263        H
const SlidingWindowSize = 3

func partTwo(nums []int) int {
	// Need at least two windows
	if len(nums) < SlidingWindowSize*2 {
		return 0
	}

	sums := []int{}
	max := len(nums) - (SlidingWindowSize - 1)

	for i := 0; i < max; i++ {
		sum := 0
		for b := 0; b < SlidingWindowSize; b++ {
			sum += nums[i+b]
		}

		sums = append(sums, sum)
	}

	return increased(sums)
}

func main() {
	fmt.Println(increased(inputs.DayOneDepths))
	fmt.Println(partTwo(inputs.DayOneDepths))
}
