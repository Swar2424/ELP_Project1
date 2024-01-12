package main

func Counting(slice []int, dijkstra [][]int, Canal2 chan (int), origin int) {
	for slice[0] != origin {
		Canal2 <- slice[0]
		slice = dijkstra[slice[0]]
	}
	Canal2 <- slice[0]
}

func Launch_Counting(graph [][]int, Canal1 chan (Table_int), Canal2 chan int) {
	for i := 0; i < len(graph); i++ {
		rep := <-Canal1
		dijkstra := rep.Table
		origin := rep.Origin
		for _, slice := range *dijkstra {
			go Counting(slice, *dijkstra, Canal2, origin)
		}
	}
}
