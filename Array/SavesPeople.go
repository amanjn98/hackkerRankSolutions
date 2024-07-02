package Array
//881. Boats to Save People
//Topics -> Two pointers
//Companies
//You are given an array people where people[i] is the weight of the ith person, and an infinite number of boats where
//each boat can carry a maximum weight of limit. Each boat carries at most two people at the same time,
//provided the sum of the weight of those people is at most limit.
//
//Return the minimum number of boats to carry every given person.
import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	minimumBoatsRequired := 0
	s := 0
	for end := len(people) - 1; end >= s; end-- {
		if people[end] == limit {
			minimumBoatsRequired++
		} else if people[end]+people[s] <= limit {
			s++
			minimumBoatsRequired++
		} else {
			minimumBoatsRequired++
		}
	}

	return minimumBoatsRequired
}
