package matrix

import "errors"

type Matrix = [][]int

func Identity(n int) *Matrix {
	matrix := make(Matrix, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		matrix[i][i] = 1
	}

	return &matrix
}

func Multiply(matrix1 *Matrix, matrix2 *Matrix) (*Matrix, error) {
	if matrix1 == nil || matrix2 == nil {
		return nil, errors.New("empty matrix")
	}
	n := len(*matrix1)
	if n < 1 {
		return nil, errors.New("unsupported 1st matrix format")
	}
	m := len((*matrix2)[0])
	if n < 1 {
		return nil, errors.New("unsupported 2nd matrix format")
	}
	if len((*matrix1)[0]) != len((*matrix2)) {
		return nil, errors.New("unsupported matrix dimentions")
	}

	result := make(Matrix, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
	}

	l := len(*matrix2)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < l; k++ {
				result[i][j] += (*matrix1)[i][k] * (*matrix2)[k][j]
			}
		}
	}

	return &result, nil
}

func Sum(matrix1 *Matrix, matrix2 *Matrix) error {
	n := len(*matrix1)
	m := len((*matrix1)[0])
	if n < 1 || n != len(*matrix2) || m != len((*matrix1)[0]) {
		return errors.New("unsupported dimentions")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			(*matrix1)[i][j] += (*matrix2)[i][j]
		}
	}

	return nil
}

func Copy(mtr *Matrix) (*Matrix, error) {
	if mtr == nil {
		return nil, errors.New("matrix should not be nil")
	}
	n := len(*mtr)
	copy := make(Matrix, n)

	for i := 0; i < n; i++ {
		copy[i] = make([]int, n)
		for j := 0; j < n; j++ {
			copy[i][j] = (*mtr)[i][j]
		}
	}

	return &copy, nil
}
