package algo

import (
	"log"

	"github.com/leonid-voroshilov/mm-entry-prog/matrix"
)

func GetAccVert(graph *matrix.Matrix) ([]int, error) {
	res, err := matrix.Copy(graph)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(*graph); i > 0; i-- {
		tmp, err := matrix.Multiply(res, graph)
		if err != nil {
			return nil, err
		}
		err = matrix.Sum(res, tmp)
		if err != nil {
			return nil, err
		}
	}
	var isSucc bool
	var result []int
	matrix.Sum(res, matrix.Identity(len(*graph)))
	for i := len(*res) - 1; i >= 0; i-- {
		isSucc = true
		for j := len(*res) - 1; j >= 0; j-- {
			if (*res)[j][i] == 0 {
				isSucc = false
				break
			}
		}
		if isSucc {
			result = append(result, i+1)
		}
	}
	return result, nil
}
