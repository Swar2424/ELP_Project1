package main

func Counting(slice []int, dijkstra [][]int, Canal2 chan (int), Canal3 chan (int), origin int) {
	for slice[0] != origin {
		Canal2 <- slice[0]
		slice = dijkstra[slice[0]]
	}
	Canal2 <- slice[0]
	Canal3 <- -1
}

func Launch_Counting(graph [][]int, Canal1 chan (Table_int), Canal2 chan int, Canal3 chan (int)) {
	n := len(graph)
	count := n * n
	for i := 0; i < n; i++ {
		rep := <-Canal1
		dijkstra := rep.Table
		//fmt.Println(dijkstra)
		origin := rep.Origin
		for _, slice := range *dijkstra {
			go Counting(slice, *dijkstra, Canal2, Canal3, origin)
		}
	}
	for count != 0 {
		count += <-Canal3
	}
	close(Canal2)
	close(Canal3)
}
