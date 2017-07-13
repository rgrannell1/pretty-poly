
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

func Geohash (precision int8, bounds interval, num float64) geohash {

	values := make([ ]bool, precision, precision)

	var isUpperBucket bool

	tmpInterval := interval {
		lower: bounds.lower,
		upper: bounds.upper,
	}

	for ith := 0; ith < int(precision); ith++ {

		offset := (bounds.upper - bounds.lower) / math.Pow(2, float64(ith + 1))
		pivot  := tmpInterval.lower + ((tmpInterval.upper - tmpInterval.lower) / 2.0)

		isUpperBucket = num > float64(pivot)

		values[ith] = isUpperBucket

		if pivot < bounds.lower || pivot > bounds.upper {
			panic("pivot out of bounds.")
		}

		if isUpperBucket {
			tmpInterval.lower += offset
		} else {
			tmpInterval.upper -= offset
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

	_ = Geohash(precision, interval.x, point2d.x).values

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
	return fromBitsLittleEndian(IntersperseBool(hash.xs, hash.ys)), nil
}

func Uint64AsGeohash2d (precision int8, hash uint64) (geohash2d, error) {

	bitCount    := 2 * int(precision)
	bits        := toBits(hash, bitCount)
	xs, ys, err := DisperseBool(bits)

	if err != nil {
		return geohash2d {xs: nil, ys: nil}, err
	} else {

		return geohash2d {
			xs: xs,
			ys: ys,
		}, nil

	}

}

func (hash geohash) AsPoint (interval interval) float64 {

	point := 0.0

	for ith := 0; ith < len(hash.values); ith++ {

		divisor := math.Pow(2, float64(ith + 1))

		if hash.values[ith] == true {
			point += (interval.upper - interval.lower) / float64(divisor)
		}

	}

	return math.Floor(point)

}

func (hash geohash2d) AsPoint (interval interval2d) point2d {

	x := (geohash {values: hash.xs }).AsPoint(interval.x)
	y := (geohash {values: hash.ys }).AsPoint(interval.y)

	return point2d {
		x: x,
		y: y,
	}

}
