package main

func Dijkstra(graph [][]int, x int) Table_int {
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
	rep := Table_int{&tab_dist, x}

	return (rep)
}

func Worker_dijkstra(graph [][]int, Jobs chan int, Results chan (Table_int)) {
	for j := range Jobs {
		Results <- Dijkstra(graph, j)
	}
}

func Launcher(n int, Jobs chan int, Results_dij chan Table_int, Results_count chan int, Count_go chan int) {
	Counter_go := n
	for i := 0; i < n; i++ {
		Jobs <- i
	}
	for Counter_go != 0 {
		Counter_go += <-Count_go
	}
	close(Count_go)
	close(Results_count)
	close(Results_dij)
	close(Jobs)
}
