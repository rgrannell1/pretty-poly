
package pretty_poly





type StaticError struct {
	string
}

func (err StaticError) Error( ) string {
	return err.string
}




var (
	ErrMisbalancedBits = StaticError { "will not divide an odd-number of bits into two slices." }
)
