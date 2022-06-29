package others

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		edges[prerequisite[0]] = append(edges[prerequisite[0]], prerequisite[1])
		indeg[prerequisite[1]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0
	for len(queue) > 0 {
		visited++
		now := queue[0]
		queue = queue[1:]
		for _, node := range edges[now] {
			indeg[node]--
			if indeg[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	return visited == numCourses
}
