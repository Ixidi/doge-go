package math

func Multiple4x4Matrices(a [16]float32, b [16]float32) (result [16]float32, err error) {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var c float32 = 0
			for k := 0; k < 4; k++ {
				c += a[(i*4)+k] * b[(k*4)+j]
			}
			result[(i*4)+j] = c
		}
	}

	return
}
