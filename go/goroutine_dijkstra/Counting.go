package main

func Counting(slice []int, dijkstra [][]int, Canal2 chan (int), origin int) {
	if slice[0] != origin {
		Counting(dijkstra[slice[0]], dijkstra, Canal2, origin)
	}
	Canal2 <- slice[0]
}

func Counters(dijkstra [][]int, Canal2 chan (int), Canal3 chan (int), origin int) {
	for _, slice := range dijkstra {
		Counting(slice, dijkstra, Canal2, origin)
	}
	Canal3 <- -1
}

func Launch_Counting(graph [][]int, Canal1 chan (Table_int), Canal2 chan int, Canal3 chan (int)) {
	n := len(graph)
	count := n
	for i := 0; i < n; i++ {
		rep := <-Canal1
		dijkstra := rep.Table
		origin := rep.Origin
		go Counters(*dijkstra, Canal2, Canal3, origin)
	}
	for count != 0 {
		count += <-Canal3

	}
	close(Canal2)
	close(Canal3)
}
