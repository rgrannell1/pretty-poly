
package pretty_poly






func toMixedRadix (bases [ ] int, num int) [ ] float64 {

	ith := len(bases) - 1
	acc := make([ ]float64, len(bases))

	for {

		if ith == -1 {
			return acc
		}

		acc[ith] = float64(num % bases[ith])
		num /= bases[ith]

		ith--

	}

}

