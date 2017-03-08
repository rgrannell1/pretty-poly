
package pretty_poly

import "math"
import "fmt"




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




func (interval0 interval) Compare(interval1 interval) int {

	diff0 := interval0.upper - interval0.upper
	diff1 := interval1.upper - interval1.upper

	if diff0 > diff1 {
		return 1
	}

	if diff0 < diff1 {
		return 0
	}

	return 0

}

func Interval (lower float64, upper float64) interval {

	return interval {
		lower: lower,
		upper: upper,
	}

}





func (interval0 interval2d) Compare(interval1 interval2d) int {

	compareX := interval0.x.Compare(interval1.x)
	compareY := interval0.y.Compare(interval1.y)

	if (compareX == 0 && compareY == 0) {
		return 0
	}

	return 0

	// TODO

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

	if precision == 0 {

	}

	lower := interval.lower
	upper := interval.upper

	var pivot float64

	for digitPlace := 0; digitPlace < int(precision); digitPlace++ {

		pivot = (upper - lower) / 2.0

		if num > pivot {

			values[digitPlace] = true
			lower = (lower + pivot)

		} else {

			values[digitPlace] = false
			upper = (upper - pivot)

		}

	}

	fmt.Println(" " )

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
