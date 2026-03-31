func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	al := make(map[int][]int)
	for _, p := range prerequisites {
		al[p[0]] = append(al[p[0]], p[1])
	}
	state := make([]int, numCourses) // 0=unvisited, 1=inStack, 2=done
	order := []int{}

	var dfs func(node int) bool // if true returned, cycle found
	dfs = func(node int) bool {
		if state[node] == 1 {return true} // still in stack, cycle found
		if state[node] == 2 {return false} // fully visited

		// like setting visited[node] to true and inStack[node] to true
		state[node] = 1
		for _, dep := range al[node] {
			// look for cycles
			if dfs(dep) {
                  return true
            }
		}
		// equivalent to setting inStack[node] to false
		state[node] = 2
		order = append(order, node)
		return false
	}

	for i:=0; i<numCourses; i++ {
		if dfs(i) {
			return []int{}
		}
	}

	return order
}