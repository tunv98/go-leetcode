package main

import (
	"fmt"
	"testing"
)

/*
graph is point and cạnh liền kề của nó
=> triển khai bằng đệ quy hoặc ngăn sếp
*/
func dfs(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] == true {
		return
	}
	visited[node] = true
	fmt.Print(node, " ")
	for _, neighbor := range graph[node] {
		dfs(graph, neighbor, visited)
	}
}

func bfs(graph map[int][]int, start int) {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func Test_DFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0},
		3: {1, 5},
		4: {1},
		5: {3, 6, 7, 8},
		6: {5},
		7: {5},
		8: {5, 9},
		9: {8},
	}
	visited := make(map[int]bool)
	fmt.Print("DFS Traversal: ")
	dfs(graph, 0, visited)
}

func Test_BFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0},
		3: {1, 5},
		4: {1},
		5: {3, 6, 7, 8},
		6: {5},
		7: {5},
		8: {5, 9},
		9: {8},
	}
	fmt.Print("BFS Traversal: ")
	bfs(graph, 0)
}
