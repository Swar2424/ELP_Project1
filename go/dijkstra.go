package main

func Dijkstra(graph [][]int, x int, Canal chan ([][]int)) {
	graph2 := Create_matrix(len(graph))
	for i := range graph {
		for j := range graph[i] {
			graph2[i][j] = graph[i][j] * x
		}
	}
	Canal <- graph2
}
