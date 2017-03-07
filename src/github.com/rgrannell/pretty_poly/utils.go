
package pretty_poly





func productOf (elems [ ]int) int {

	prod := 1

	for _, elem := range elems {
		prod = prod * elem
	}

	return prod

}
