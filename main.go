package main

import (
	"fmt"
	"math"
)

func minSteps(step []int, sum int) int {

	if sum == 0 {
		return 0
	}
	res := math.MaxInt32

	for _, val := range step {
		if val <= sum {
			sub_res := minSteps(step, sum-val)
			if sub_res != math.MaxInt32 && sub_res+1 < res {
				res = sub_res + 1
			}
		}

	}
	return res
}
func noOfWays(step []int,sum int)int {
	if sum == 0 {
		return 1
	}
	if sum < 0 {
		return 0
	}
	if len(step) == 0 && sum >= 1{
		return 0
	}
	return noOfWays(step,sum - step[len(step)-1]) + noOfWays(step[:len(step)-1],sum)
}
func main() {
	var dest int
	steps := []int{3, 4, 5}
	fmt.Scanln(&dest)
	fmt.Println(minSteps(steps, dest))
	fmt.Println(noOfWays(steps,dest))
}

