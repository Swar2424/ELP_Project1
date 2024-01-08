package main

func Counting(slice []int, dijkstra [][]int, Canal2 chan (int), Canal3 chan (int), origin int) {
	if slice[0] != origin {
		Canal3 <- 1
		go Counting(dijkstra[slice[0]], dijkstra, Canal2, Canal3, origin)
	}
	Canal2 <- slice[0]
	Canal3 <- -1
}

func Launch_Counting(graph [][]int, Canal1 chan ([][]int), Canal2 chan int, Canal3 chan int) {
	count := 0

	for i := 0; i < len(graph); i++ {
		dijkstra := <-Canal1
		origin := -1
		for _, slice := range dijkstra {
			if slice[1] == 0 {
				origin = slice[0]
			}
		}
		for _, slice := range dijkstra {
			count += 1
			go Counting(slice, dijkstra, Canal2, Canal3, origin)
		}
	}
	for count != 0 {
		count += <-Canal3
	}
	close(Canal2)

}
