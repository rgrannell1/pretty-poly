
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






func Geohash2dAsUint64 (hash geohash2d) uint64 {

	storable := uint64(0)

	for ith := 0; ith < len(hash.xs); ith++ {
		for jth := 0; jth < 2; jth++ {

			var bisection bool

			if jth == 0 {
				bisection = hash.xs[ith]
			} else {
				bisection = hash.ys[ith]
			}

			if bisection {
				storable += uint64(2^ith)
			}

		}
	}

	return storable

}





func uint64AsGeohash2d (precision int8, hash uint64) geohash2d {

	digits := int8(math.Ceil( math.Log2(float64(hash)) ))

	xs := [ ] bool{ }
	ys := [ ] bool{ }

	var sw int

	if (digits < precision) {

		sw = 0

		for ith := int8(0); ith < (precision - int8(digits)); ith++ {

			if sw % 2 == 0 {
				xs = append(xs, false)
			} else {
				ys = append(ys, false)
			}

			sw++

		}

	}

	sw = 0

	for ith := (digits - 1); ith >= 0; ith++ {

	}

	return geohash2d {
		xs: xs,
		ys: ys,
	}

}
