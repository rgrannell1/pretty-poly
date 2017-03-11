
package pretty_poly





type StaticError struct {
	string
}

func (err StaticError) Error( ) string {
	return err.string
}




var (
	ErrMisbalancedGeohash = StaticError { "foo!" }
)
