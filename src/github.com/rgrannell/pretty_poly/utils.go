
package pretty_poly




//import "fmt"
import "math"





func productOf (elems [ ]int) int {

	prod := 1

	for _, elem := range elems {
		prod = prod * elem
	}

	return prod

}

func toBits (hash uint64, length int) [ ] bool {

	if hash == 0 {
		return make([ ] bool, length, length)
	}

	digits := int(math.Floor(math.Log2(float64(hash)) )) + 1
	bits   := make( [ ] bool, length, length)

	for ith := 0; ith < length; ith++ {
		bits[ith] = false
	}

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

func IntersperseBool (bits0, bits1 [ ] bool) [ ] bool {

	outSize := len(bits0) + len(bits1)
	output  := make([ ] bool, outSize, outSize)

	bits0Count := 0
	bits1Count := 0

	for ith := 0; ith < outSize; ith += 2 {

		output[ith] = bits0[bits0Count]
		bits0Count++

	}

	for ith := 1; ith < outSize; ith += 2 {

		output[ith] = bits1[bits1Count]
		bits1Count++

	}

	return output

}

func DisperseBool (bits [ ] bool) ([ ] bool, [ ] bool, error) {

	if len(bits) % 2 != 0 {
		return nil, nil, ErrMisbalancedBits
	}

	xs := make([ ] bool, len(bits) / 2, len(bits) / 2)
	ys := make([ ] bool, len(bits) / 2, len(bits) / 2)

	xcount := 0
	ycount := 0

	for ith := 0; ith < len(bits); ith++ {

		if ith % 2 == 0 {

			xs[xcount] = bits[ith]
			xcount++

		} else {

			ys[ycount] = bits[ith]
			ycount++

		}

	}

	return xs, ys, nil

}
