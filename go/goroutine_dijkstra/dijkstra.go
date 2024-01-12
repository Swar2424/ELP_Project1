package main

func Dijkstra(graph [][]int, x int, Canal chan (*[][]int)) { //on ne prend pas en compte le décalage
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
					if graph[i][j] != 0 && tab_visit[j] == 1 {
						dist := graph[i][j] + tab_dist[j][1]
						if dist < tab_dist[i][1] {
							tab_dist[i] = []int{j, dist}
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
	Canal <- &tab_dist
}

func Launch_Dijkstra(graph [][]int, Canal chan (*[][]int)) {
	for i := 0; i < len(graph); i++ {
		go Dijkstra(graph, i, Canal)
	}
}
