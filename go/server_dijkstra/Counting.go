package main

func Counting(slice []int, dijkstra [][]int, Results chan (int), origin int) {
	if slice[0] != origin {
		Counting(dijkstra[slice[0]], dijkstra, Results, origin)
	}
	Results <- slice[0]
}

func Counters(dijkstra [][]int, Results chan (int), Count_go chan (int), origin int) {
	for _, slice := range dijkstra {
		Counting(slice, dijkstra, Results, origin)
	}
	Count_go <- -1
}

func Worker_counter(Jobs chan (Table_int), Results chan int, Count_go chan int) {
	for j := range Jobs {
		Counters(*j.Table, Results, Count_go, j.Origin)
	}
}
