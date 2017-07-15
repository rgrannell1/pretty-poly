
package pretty_poly




func validateSolveArguments (order int, extreme int, filename string) error {

	if order <= 0 {
		return ErrOrderArgumentSize
	}

	return nil

}

func Solve (order int, extreme int, precision int8, filename string) error {

	err := validateSolveArguments(order, extreme, filename)

	if (err != nil) {

		return err

	} else {

		SolvePolynomials(extreme, order, filename, precision)
		return nil

	}

}





func validateDrawArguments (filename string, precision int8) error {
	return nil
}

func Draw (filename string, precision int8) error {
	return DrawImage(filename, float64(precision))
}
