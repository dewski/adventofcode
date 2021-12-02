package main

import (
	"fmt"

	"github.com/dewski/adventofcode/2021/inputs"
)

func main() {
	fmt.Printf("Part One: %d\n", increased(inputs.DayOneDepths))
	fmt.Printf("Part Two: %d\n", partTwo(inputs.DayOneDepths))
}

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
	// Need at least 2 windows of numbers to be able to compare increase
	if len(nums) < SlidingWindowSize*2 {
		return 0
	}

	sums := []int{}
	max := len(nums) - (SlidingWindowSize - 1)

	for i := 0; i < max; i++ {
		sum := 0
		// Lookup next N (SlidingWindowSize) elements of provided numbers from
		// current index to calculate window size.
		//
		// 199  i=0, b=0
		// 200  i=0, b=1
		// 208  i=0, b=2
		// 199 + 200 + 208 = 607
		for b := 0; b < SlidingWindowSize; b++ {
			sum += nums[i+b]
		}

		sums = append(sums, sum)
	}

	return increased(sums)
}
