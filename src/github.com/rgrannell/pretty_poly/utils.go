
package pretty_poly




import "math"





func productOf (elems [ ]int) int {

	prod := 1

	for _, elem := range elems {
		prod = prod * elem
	}

	return prod

}

func toBits (hash uint64) [ ] bool {

	digits := math.Ceil( math.Log2(float64(hash)) )


	bits := make( [ ] bool, int(digits))

	for {
		if hash <= 0 {
			break
		}

		bits = append(bits, hash % 2 == 1)
		hash /= 2

	}

	return bits

}
