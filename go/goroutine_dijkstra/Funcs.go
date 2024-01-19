package main

func Create_matrix(N int) [][]int {
	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, N)
	}
	return a
}

type Table_int struct {
	Table  *[][]int
	Origin int
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
