package q2

const l = 256

func GeMM1(A [][l]uint, B [][l]uint) (C [l][l]uint) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			for k := 0; k < len(B); k++ { // or len(A[0])
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return
}

func GeMM2(A [][l]uint, B [][l]uint) (C [l][l]uint) { // assuming B is n × n for code simplicity
	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B[0]); j++ {
			B[i][j], B[j][i] = B[j][i], B[i][j]
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			for k := 0; k < len(B[0]); k++ { // or len(A[0])
				C[i][j] += A[i][k] * B[j][k]
			}
		}
	}
	return
}
