package Graph

import "container/heap"

//You are given an integer n indicating there are n people numbered from 0 to n - 1.
//You are also given a 0-indexed 2D integer array meetings where meetings[i] = [xi, yi, timei] indicates that person xi
//and person yi have a meeting at time i. A person may attend multiple meetings at the same time.
//Finally, you are given an integer firstPerson.
// Person 0 has a secret and initially shares the secret with a person firstPerson at time 0.
//This secret is then shared every time a meeting takes place with a person that has the secret.
//More formally, for every meeting, if a person xi has the secret at timei, then they will share the secret with person yi,
//and vice versa.
// The secrets are shared instantaneously. That is, a person may receive the secret and share it with people in other meetings
//within the same time frame.
//
//Return a list of all the people that have the secret after all the meetings have taken place. You may return the answer in any order.
//Example 1:
//
//Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1
//Output: [0,1,2,3,5]
//Explanation:
//At time 0, person 0 shares the secret with person 1.
//At time 5, person 1 shares the secret with person 2.
//At time 8, person 2 shares the secret with person 3.
//At time 10, person 1 shares the secret with person 5.
//Thus, people 0, 1, 2, 3, and 5 know the secret after all the meetings.
//Example 2:
//
//Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3
//Output: [0,1,3]
//Explanation:
//At time 0, person 0 shares the secret with person 3.
//At time 2, neither person 1 nor person 2 know the secret.
//At time 3, person 3 shares the secret with person 0 and person 1.
//Thus, people 0, 1, and 3 know the secret after all the meetings.
//Example 3:
//
//Input: n = 5, meetings = [[3,4,2],[1,2,1],[2,3,1]], firstPerson = 1
//Output: [0,1,2,3,4]
//Explanation:
//At time 0, person 0 shares the secret with person 1.
//At time 1, person 1 shares the secret with person 2, and person 2 shares the secret with person 3.
//Note that person 2 can share the secret at the same time as receiving it.
//At time 2, person 3 shares the secret with person 4.
//Thus, people 0, 1, 2, 3, and 4 know the secret after all the meetings.

//Approach
//Build a Graph: Construct a graph where nodes represent people and edges represent meetings, weighted by their time.
//Employ BFS with Priority Queue: Utilize a priority queue to efficiently traverse the graph in a breadth-first manner,
//prioritizing earlier meetings.
//Ensure Time-Based Propagation: Only propagate the secret to a person if the meeting time is later than or equal to the
//time the secret was acquired.

type MinHeap []Meet

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Meet))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Meet struct {
	person int
	time   int
}

func FindAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make([][]Meet, n)
	for _, meeting := range meetings {
		graph[meeting[0]] = append(graph[meeting[0]], Meet{meeting[1], meeting[2]})
		graph[meeting[1]] = append(graph[meeting[1]], Meet{meeting[0], meeting[2]})
	}
	visited := make([]bool, n)
	queue := MinHeap{{0, 0}, {firstPerson, 0}}

	for len(queue) > 0 {
		f := heap.Pop(&queue)
		if visited[f.(Meet).person] {
			continue
		}
		visited[f.(Meet).person] = true
		for _, v := range graph[f.(Meet).person] {
			if v.time >= f.(Meet).time && !visited[v.person] {
				heap.Push(&queue, v)
			}
		}
	}

	ans := []int{}
	for i, v := range visited {
		if v {
			ans = append(ans, i)
		}
	}
	return ans
}
