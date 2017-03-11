
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

	if hash == 0 {
		return [ ] bool { }
	}

	digits := int(math.Floor(math.Log2(float64(hash)) )) + 1
	bits   := make( [ ] bool, digits)

	for {
		if hash <= uint64(0) {
			break
		}

		bits[digits-1] = hash % 2 == 1
		digits--

		hash /= 2

	}

	return bits

}

func fromBitsLittleEndian (bits [ ] bool) uint64 {

	output := uint64(0)

	for ith := 0; ith < len(bits); ith++ {

		if bits[ith] {
			output += uint64( math.Exp2(float64(ith)) )
		}

	}

	return output

}

func intersperse (bits0 [ ] bool , bits1 [ ] bool) [ ] bool {

	output := make([ ] bool, len(bits0) + len(bits1), len(bits0) + len(bits1))


	return output

}
