
package pretty_poly

import "math"




type geohash struct {
	values [ ] bool
}

type geohash2d struct {
	xs [ ] bool
	ys [ ] bool
}

type interval struct {
	lower float64
	upper float64
}

type interval2d struct {
	x interval
	y interval
}

type point2d struct {
	x float64
	y float64
}





func Interval (lower float64, upper float64) interval {

	return interval {
		lower: lower,
		upper: upper,
	}

}

func (interval0 *interval) AddXInterval (interval1 interval) interval2d {
	return interval2d {
		x: interval1,
		y: *interval0,
	}
}

func (interval0 *interval) AddYInterval (interval1 interval) interval2d {
	return interval2d {
		x: *interval0,
		y: interval1,
	}
}

func Interval2d (lowerx float64, upperx float64, lowery float64, uppery float64) interval2d {

	return interval2d {
		x: Interval(lowerx, upperx),
		y: Interval(lowery, uppery),
	}

}






func (hash0 geohash) Equal(hash1 geohash) bool {

	for ith := 0; ith < len(hash0.values); ith++ {
		if hash0.values[ith] != hash1.values[ith] {
			return false
		}
	}

	return true

}

func (hash geohash) Decompress ( ) geohash2d {

	xs := make([ ] bool, len(hash.values) / 2, len(hash.values) / 2)
	ys := make([ ] bool, len(hash.values) / 2, len(hash.values) / 2)

	xcount := 0
	ycount := 0

	for ith := 0; ith < len(hash.values); ith++ {

		if ith % 2 == 0 {

			xs[xcount] = hash.values[ith]
			xcount++

		} else {

			ys[ycount] = hash.values[ith]
			ycount++

		}

	}

	return geohash2d {
		xs: xs,
		ys: ys,
	}

}

func Geohash (precision int8, interval interval, num float64) geohash {

	values := make([ ]bool, precision, precision)

	var isUpperBucket bool
	var pivot float64

	for ith := 0; ith < int(precision); ith++ {

		pivot         = interval.lower + ((interval.upper - interval.lower) / 2.0)
		isUpperBucket = num > pivot

		values[ith] = isUpperBucket

		if isUpperBucket {
			interval.lower += pivot
		} else {
			interval.upper -= pivot
		}

	}

	return geohash {
		values: values,
	}

}






func (hash0 geohash2d) Equal(hash1 geohash2d) bool {

	for ith := 0; ith < len(hash0.xs); ith++ {
		if hash0.xs[ith] != hash1.xs[ith] {
			return false
		}
	}

	for ith := 0; ith < len(hash0.ys); ith++ {
		if hash0.ys[ith] != hash1.ys[ith] {
			return false
		}
	}

	return true

}





func (hash geohash2d) Compress ( ) geohash {

	outputLength := len(hash.xs) + len(hash.ys)
	output       := make([ ] bool, outputLength, outputLength)

	xcount := 0
	ycount := 0

	for ith := 0; ith < outputLength; ith++ {

		if ith % 2 == 0 {

			output[ith] = hash.xs[xcount]
			xcount++

		} else {

			output[ith] = hash.ys[ycount]
			ycount++

		}

	}

	return geohash {
		values: output,
	}

}





func Geohash2d (precision int8, interval interval2d, point2d point2d) geohash2d {

	return geohash2d {
		xs: Geohash(precision, interval.x, point2d.x).values,
		ys: Geohash(precision, interval.y, point2d.y).values,
	}

}

func (hash0 *geohash) AddXAxis (hash1 geohash) geohash2d {
	return geohash2d {
		xs: hash1.values,
		ys: hash0.values,
	}
}

func (hash0 *geohash) AddYAxis (hash1 geohash) geohash2d {
	return geohash2d {
		xs: hash0.values,
		ys: hash1.values,
	}
}






func Geohash2dAsUint64 (hash geohash2d) (uint64, error) {

	storable := uint64(1)

	if len(hash.xs) != len(hash.ys) {

		return 0, ErrMisbalancedGeohash

	} else if len(hash.xs) == 0 {

		return 0, nil

	} else {

		digit := 0

		for ith := 0; ith < len(hash.xs); ith++ {

			if hash.xs[ith] {

				storable += uint64(2^digit)
				digit++

			}

			if hash.ys[ith] {

				storable += uint64(2^digit)
				digit++

			}

		}

		return storable, nil

	}

}





func Uint64AsGeohash2d (precision int8, hash uint64) geohash2d {

	digits := int8(math.Ceil( math.Log2(float64(hash)) ))

	xs := make([ ] bool, precision, precision)
	ys := make([ ] bool, precision, precision)

	// append leading zeros.
	if (digits < precision) {

		for ith := int8(0); ith < (precision - int8(digits)); ith++ {

			if ith % 2 == 0 {
				xs[ith] = false
			} else {
				ys[ith] = false
			}

		}

	}

	for ith := (digits - 1); ith >= 0; ith++ {

	}

	return geohash2d {
		xs: xs,
		ys: ys,
	}

}
