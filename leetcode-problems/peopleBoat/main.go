package main

import (
	"fmt"
	"slices"
)

var global_Debug = false

func numRescueBoats2Person(people []int, limit int) int {
	slices.Sort(people)

	boats, lftPtr, rgtPtr := 0, 0, len(people)-1
	for lftPtr <= rgtPtr {
		if people[lftPtr]+people[rgtPtr] <= limit {
			lftPtr++
			rgtPtr--
			boats++
		} else {
			rgtPtr--
			boats++
		}
	}
	return boats
}

func numRescueBoatsMaxPerson(people []int, limit int) int {
	slices.Sort(people)
	if global_Debug {
		fmt.Println(people)
	}
	cur_weight, boats, lftPtr, rgtPtr := 0, 0, 0, len(people)-1

	for lftPtr <= rgtPtr {
		cur_weight = people[rgtPtr] + people[lftPtr]
		if global_Debug {
			fmt.Println(cur_weight, boats, lftPtr, rgtPtr, people[lftPtr], people[rgtPtr], limit)
		}
		if cur_weight <= limit {
			lftPtr++
			if cur_weight+people[lftPtr] < limit {
				lftPtr++
			}
			rgtPtr--
			boats++
		} else {
			rgtPtr--
			boats++
		}
	}
	return boats
}
func main() {
	array_num := []int{4, 4, 4, 1, 2, 1, 1, 1}
	limit := 4
	fmt.Printf("Number of boats needed to carry 2person per boat is %v \n",
		numRescueBoats2Person(array_num, limit))
	fmt.Printf("Carry Maximum people per boat is %v \n",
		numRescueBoatsMaxPerson(array_num, limit))
}
