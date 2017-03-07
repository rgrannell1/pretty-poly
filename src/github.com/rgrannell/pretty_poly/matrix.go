
package pretty_poly



import "github.com/gonum/matrix/mat64"





func toCompanionMatrix (coeffs [ ] float64) mat64.Dense {

	matrix := mat64.NewDense(len(coeffs), len(coeffs), nil)

	for ith := 0; ith < len(coeffs); ith++ {

		if ith > 0 {
			matrix.Set(ith, ith - 1, 1)
		}

		matrix.Set(ith, len(coeffs) - 1, -coeffs[ith])

	}

	return *matrix

}
