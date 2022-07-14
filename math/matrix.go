package math

import "errors"

func Multiple4x4Matrices(a []float64, b []float64) ([]float64, error) {
	if len(a) != 16 || len(a) != len(b) {
		return nil, errors.New("4x4 matrices were expected")
	}

	result := make([]float64, 16)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c := .0
			for k := 0; k < 4; k++ {
				c += a[(i*4)+k] * b[(k*4)+j]
			}
			result[(i*4)+j] = c
		}
	}

	return result, nil
}
