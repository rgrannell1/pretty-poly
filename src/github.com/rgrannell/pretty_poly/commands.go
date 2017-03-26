
package pretty_poly




func validateSolveArguments (order int, extreme int, filename string) error {

	if order <= 0 {
		return ErrOrderArgumentSize
	}

	return nil

}

func Solve (order int, extreme int, filename string) error {

	err := validateSolveArguments(order, extreme, filename)

	if (err != nil) {

		return err

	} else {

		SolvePolynomials(extreme, order, filename)
		return nil

	}

}





func validateDrawArguments (filename string) error {
	return nil
}

func Draw (filename string) {
	DrawImage(filename)
}
