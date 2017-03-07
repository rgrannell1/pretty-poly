
package pretty_poly





import "testing"
import "github.com/franela/goblin"




type geoHashTestCase struct {
	precision int8
	interval  interval
	num       float64
	result    [ ] bool
}





func BenchmarkGeohash (bench *testing.B) {

	interval := Interval(0, 60000)

	var _ geohash

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		_ = Geohash(8, interval, 30000)
	}

}





func BenchmarkGeohash2d (bench *testing.B) {

	interval := Interval2d(0, 60000, 0, 60000)
	point    := point {
		x: 30000,
		y: 30000,
	}

	var _ geohash

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		_ = Geohash2d(8, interval, point)
	}

}




func BenchmarkGeohash2dAsUint64 (bench *testing.B) {

	interval := Interval2d(0, 60000, 0, 60000)
	point    := point {
		x: 30000,
		y: 30000,
	}

	geohash := Geohash2d(8, interval, point)

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		_ = Geohash2dAsUint64(geohash)
	}

}





func BenchmarkUint64AsGeohash2d (bench *testing.B) {

	interval := Interval2d(0, 60000, 0, 60000)
	point    := point {
		x: 30000,
		y: 30000,
	}

	geohash := Geohash2d(8, interval, point)
	hashInt := Geohash2dAsUint64(geohash)

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		uint64AsGeohash2d(8, hashInt)
	}

}





func TestGeohashCreation (test *testing.T) {

	gob       := goblin.Goblin(test)
	testCases := [ ] geoHashTestCase {

		geoHashTestCase {
			precision: 1,
			interval:  Interval(0, 2),
			num:       0,
			result:   [ ] bool {false},
		},
		geoHashTestCase {
			precision: 1,
			interval:  Interval(0, 2),
			num:       2,
			result:   [ ] bool {true},
		},
		geoHashTestCase {
			precision: 5,
			interval:  Interval(0, 100),
			num:       0,
			result:   [ ] bool {false, false, false, false, false},
		},
		geoHashTestCase {
			precision: 2,
			interval:  Interval(0, 100),
			num:       51,
			result:   [ ] bool {true, false},
		},
	}

	gob.Describe("Geohash", func ( ) {

		var result geohash

		for _, testCase := range testCases {

			result = Geohash(testCase.precision, testCase.interval, testCase.num)

			gob.It("generates the expected output value", func ( ) {

				for ith := 0; ith < len(testCase.result); ith++ {
					result = result
					gob.Assert(testCase.result[ith]).Equal(result.values[ith])
				}

			})

		}

	})

	gob.Describe("Geohash2d", func ( ) {

		var result         geohash2d
		var testInterval2d interval2d
		var testPoint      point

		for _, testCase := range testCases {

			testInterval2d = interval2d {
				x: testCase.interval,
				y: testCase.interval,
			}

			testPoint = point {
				x: testCase.num,
				y: testCase.num,
			}

			result = Geohash2d(testCase.precision, testInterval2d, testPoint)

			gob.It("generates the expected output value", func ( ) {

				for ith := 0; ith < len(testCase.result); ith++ {
					result = result
					gob.Assert(testCase.result[ith]).Equal(result.xs[ith])
					gob.Assert(testCase.result[ith]).Equal(result.ys[ith])
				}

			})

		}

	})

}
