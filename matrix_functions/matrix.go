//package main
package matrix_transformations

import (
	"errors"
)

func dot_product(x []float64, y []float64) (float64, error) {
	if (len(x) != len(y)) {
		return 0, errors.New("Mismatch dimensions")
	}

	sum := float64(0)

	for i := 0 ; i < len(x); i++ {
		sum += x[i] * y[i]
	}
	return sum, nil
}

func matrix_mult(x [][]float64, y [][]float64) ([][]float64, error) {
	if (len(x[0]) != len(y)) {
		return nil, errors.New("Mismatch dimensions")
	}

	result := make([][]float64, len(x))

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			for k := 0; k < len(y); k++ {
				result[i][j] += x[i][k] * y[k][j]
			} 
		}
	}

	return result, nil

}

func matrix_transpose(x [][]float64) ([][]float64, error) {
	result := make([][]float64, len(x), len(x[0]))

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			result[j][i] = x[i][j]
		}
	}

	return result, nil
}

func matrix_vector_mult(x [][]float64, v []float64) ([]float64, error) {
	result, err := vector_mult(v, x)
	if err != nil {
		return result, err
	}
	return result, err
}

func vector_mult(x []float64, y [][]float64) ([]float64, error) {
	if (len(x) != len(y)) {
		return nil, nil
	}

	result := make([]float64, len(y[0]))

	for i := 0; i < len(y); i++ {
		for j := 0; j < len(x); j++ {
			result[j] += y[j][i] * x[j]
		}
	}

	return result, nil
}

// func main() {
// 	x := []float64 {1, 2, 3}
// 	y := []float64 {1, 2, 3}
// 	fmt.Println(dot_product(x, y))

// 	x = x[0:len(x) -1]
// 	fmt.Println(dot_product(x, y))
// }
