package Graph


// Kahn's algorithm or BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	// Store the indegree of each course
	indegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		indegree[pre[0]]++
	}

	// Store the courses that can be completed (indegree == 0) in queue and in result array
	var queue, completedCourses []int
	for course, degree := range indegree {
		if degree == 0 {
			queue = append(queue, course)
			completedCourses = append(completedCourses, course)
		}
	}

	for len(queue) != 0 {

		// This course is completed
		course := queue[0]
		queue = queue[1:]

		for _, pre := range prerequisites {

			// If course is the prerequite of any other course?
			if course == pre[1] {

				// If yes, then reduce the indegree of that course
				indegree[pre[0]]--

				// Is there a cycle?
				if indegree[pre[0]] < 0 {
					return []int{}
				}

				// Can this course be completed now?
				if indegree[pre[0]] == 0 {
					completedCourses = append(completedCourses, pre[0])
					queue = append(queue, pre[0])
				}
			}
		}
	}

	if len(completedCourses) == numCourses {
		return completedCourses
	}

	return []int{}
}