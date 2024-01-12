package main

func Dijkstra(graph [][]int, x int) [][]int { //on ne prend pas en compte le décalage
	N := len(graph[0])
	tab_visit := make([]int, N)
	tab_dist := make([][]int, N)
	fin := 0

	for i := 0; i < N; i++ {
		if i == x {
			tab_dist[i] = []int{i, 0}
			tab_visit[i] = 1
		} else {
			tab_dist[i] = []int{i, 90000}
		}
	}
	for fin < 1 {
		for i := 0; i < N; i++ {
			if tab_visit[i] == 0 {
				for j := 0; j < N; j++ {
					//fmt.Println(graph[i][j], tab_visit[j])
					if graph[i][j] != 0 && tab_visit[j] == 1 {
						dist := graph[i][j] + tab_dist[j][1]
						if dist < tab_dist[i][1] {
							tab_dist[i] = []int{j, dist}
						}
						tab_visit[i] = 1
					}
				}
			}
			//fmt.Println(x, i, tab_visit[i], tab_dist[i])
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
