package main

import (
	"fmt"
	"log"

	"github.com/leonid-voroshilov/mm-entry-prog/algo"
	"github.com/leonid-voroshilov/mm-entry-prog/matrix"
)

func main() {
	graph := matrix.Matrix{
		{0, 1, 1, 0, 1},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
	result, err := algo.GetAccVert(&graph)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вершины, достижимые из любых других вершин: %v", result)
}
