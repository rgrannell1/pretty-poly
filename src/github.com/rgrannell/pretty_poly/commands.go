
package pretty_poly




func validateSolveArguments (order int, extreme int, filename string) error {

	if order <= 0 {
		return ErrOrderArgumentSize
	}

	return nil

}

func Solve (order int, extreme int, filename string) {

	err := validateSolveArguments(order, extreme, filename)

	if (err != nil) {
		panic(err)
	}

	SolvePolynomials(extreme, order, filename)

}





func validateDrawArguments (filename string) error {
	return nil
}

func Draw (filename string) {
	DrawImage(filename)
}
