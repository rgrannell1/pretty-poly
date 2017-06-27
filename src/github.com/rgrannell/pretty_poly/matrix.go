
package pretty_poly



import "github.com/gonum/matrix/mat64"





func toCompanionMatrix (coeffs [ ] float64, offset float64) *mat64.Dense {

	matrix := mat64.NewDense(len(coeffs), len(coeffs), nil)

	for ith := 0; ith < len(coeffs); ith++ {

		if ith > 0 {
			matrix.Set(ith, ith - 1, 1)
		}

		matrix.Set(ith, len(coeffs) - 1, -(coeffs[ith] - offset))

	}

	return matrix

}

func solvePolynomial (matrix *mat64.Dense) [ ] complex128 {

	var eigen mat64.Eigen
	eigen.Factorize(matrix, false, false)

	return eigen.Values(nil)

}
