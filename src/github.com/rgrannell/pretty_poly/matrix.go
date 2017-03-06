
package pretty_poly




import "github.com/gonum/matrix/mat64"





func toCompanionMatrix (coeffs [ ] float64) mat64.Dense {

	matrix := mat64.NewDense(len(coeffs), len(coeffs), nil)

	for ith := 1; ith < len(coeffs); ith++ {
		matrix.Set(ith, ith - 1, 1)
	}

	for ith := 1; ith < len(coeffs); ith++ {
		matrix.Set(len(coeffs) - 1, ith, -coeffs[ith])
	}

	return *matrix

}
