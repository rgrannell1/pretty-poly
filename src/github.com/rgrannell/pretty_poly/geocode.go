
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

type point struct {
	x float64
	y float64
}






func Interval (lower float64, upper float64) interval {

	return interval {
		lower: lower,
		upper: upper,
	}

}

func Interval2d (lowerx float64, upperx float64, lowery float64, uppery float64) interval2d {

	return interval2d {
		x: Interval(lowerx, upperx),
		y: Interval(lowery, uppery),
	}

}

func Geohash (precision int8, interval interval, num float64) geohash {

	values := [ ]bool { }

	if precision == 0 {

	}

	trackingInterval := Interval(interval.lower, interval.upper)

	var pivot float64

	for digitPlace := 0; digitPlace < int(precision); digitPlace++ {

		pivot = float64(trackingInterval.lower) / float64(trackingInterval.upper)

		if num > pivot {

			values = append(values, true)
			trackingInterval.lower = pivot

		} else {

			values = append(values, false)
			trackingInterval.upper = pivot

		}

	}

	return geohash {
		values: values,
	}

}



func Geohash2d (precision int8, interval interval2d, point point) geohash2d {

	return geohash2d {
		xs: Geohash(precision, interval.x, point.x).values,
		ys: Geohash(precision, interval.y, point.y).values,
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
