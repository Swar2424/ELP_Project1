package main

func Counting(slice []int, dijkstra [][]int, origin int, table []int) []int {
	if slice[0] != origin {
		table = Counting(dijkstra[slice[0]], dijkstra, origin, table)
	}
	table[slice[0]] += 1
	return (table)
}
