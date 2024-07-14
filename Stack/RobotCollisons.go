package Stack

import (
	"sort"
)

//There are n 1-indexed robots, each having a position on a line, health, and movement direction.
//
//You are given 0-indexed integer arrays positions, healths, and a string directions (directions[i] is either 'L' for left or 'R' for right). All integers in positions are unique.
//
//All robots start moving on the line simultaneously at the same speed in their given directions. If two robots ever share the same position while moving, they will collide.
//
//If two robots collide, the robot with lower health is removed from the line, and the health of the other robot decreases by one. The surviving robot continues in the same direction it was going. If both robots have the same health, they are both removed from the line.
//
//Your task is to determine the health of the robots that survive the collisions, in the same order that the robots were given, i.e. final heath of robot 1 (if survived), final health of robot 2 (if survived), and so on. If there are no survivors, return an empty array.
//
//Return an array containing the health of the remaining robots (in the order they were given in the input), after no further collisions can occur.
//
//Note: The positions may be unsorted.

//Example 1:
//
//
//
//Input: positions = [5,4,3,2,1], healths = [2,17,9,15,10], directions = "RRRRR"
//Output: [2,17,9,15,10]
//Explanation: No collision occurs in this example, since all robots are moving in the same direction.
//So, the health of the robots in order from the first robot is returned, [2, 17, 9, 15, 10].
//Example 2:
//
//
//
//Input: positions = [3,5,2,6], healths = [10,10,15,12], directions = "RLRL"
//Output: [14]
//Explanation: There are 2 collisions in this example. Firstly, robot 1 and robot 2 will collide, and
//since both have the same health, they will be removed from the line. Next, robot 3 and robot 4 will collide and
//since robot 4's health is smaller, it gets removed, and robot 3's health becomes 15 - 1 = 14.
//Only robot 3 remains, so we return [14].
//Example 3:
//
//
//
//Input: positions = [1,2,5,6], healths = [10,10,11,11], directions = "RLRL"
//Output: []
//Explanation: Robot 1 and robot 2 will collide and since both have the same health, they are both removed.
//Robot 3 and 4 will collide and since both have the same health, they are both removed. So, we return an empty array, [].

func SurvivedRobotsHealths(positions []int, healths []int, directions string) []int {
	robots := make(map[int][]int)
	for i, position := range positions {
		direction := 1
		if directions[i] == 'L' {
			direction = -1
		}
		robots[position] = []int{healths[i], direction, i}
	}

	sortedPositions := append([]int{}, positions...)
	sort.Ints(sortedPositions)
	stack := [][]int{}
	for _, position := range sortedPositions {
		stack = append(stack, robots[position])
		i := len(stack) - 1
		for i-1 >= 0 && stack[i][1] == -1 && stack[i-1][1] == 1 {
			right := stack[i-1]
			left := stack[i]
			survive := []int{}
			stack = stack[:i-1]
			if left[0] > right[0] {
				survive = left
			} else if left[0] < right[0] {
				survive = right
			} else {
				i -= 2
				continue
			}
			survive[0]--
			stack = append(stack, survive)
			i--
		}
	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i][2] < stack[j][2]
	})

	result := []int{}
	for _, robot := range stack {
		result = append(result, robot[0])
	}
	return result
}
