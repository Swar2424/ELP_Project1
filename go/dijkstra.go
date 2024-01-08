package main

func Dijkstra(graph [4][4]int, x int) []int { //on ne prend pas en compte le décalage
	N := len(graph[0])
	tab_visit := make([]int, N)
	tab_dist := make([]int, N)
	fin := 0

	for i := 0; i < N; i++ {
		if i == x {
			tab_dist[i] = 0
			tab_visit[i] = 1
		} else {
			tab_dist[i] = 90000
		}
	}
	for fin < 1 {
		for i := 0; i < N; i++ {
			if tab_visit[i] == 0 {
				for j := 0; j < N; j++ {
					if graph[i][j] != 1 && graph[i][j] != 0 {
						dist := graph[i][j] + tab_dist[j]
						if dist < tab_dist[i] {
							tab_dist[i] = dist
						}
						tab_visit[i] = 1
					}
				}
			}
		}
		fin = 1
		for i := 0; i < N; i++ {
			if tab_visit[i] == 0 {
				fin = 0
			}
		}
	}
	return (tab_dist)
}
