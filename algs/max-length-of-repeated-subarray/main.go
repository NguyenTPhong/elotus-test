package main

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	count := make([]int, n)
	dist := make([]int, n)
	var dfs1 func(u, p int)
	dfs1 = func(u, p int) {
		count[u] = 1
		for _, v := range graph[u] {
			if v != p {
				dfs1(v, u)
				count[u] += count[v]
				dist[u] += dist[v] + count[v]
			}
		}
	}
	dfs1(0, -1)
	var dfs2 func(u, p int)
	dfs2 = func(u, p int) {
		for _, v := range graph[u] {
			if v != p {
				dist[v] = dist[u] - count[v] + n - count[v]
				dfs2(v, u)
			}
		}
	}
	dfs2(0, -1)
	return dist
}

func main() {
	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
	fmt.Println(sumOfDistancesInTree(2, [][]int{{1, 0}}))
}
