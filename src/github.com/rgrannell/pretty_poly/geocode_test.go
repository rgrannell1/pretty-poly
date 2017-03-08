
package pretty_poly

import "fmt"




import "testing"
import "github.com/franela/goblin"




type geoHashTestCase struct {
	precision int8
	interval  interval
	num       float64
	result    geohash
}




func getTestCases ( ) [ ] geoHashTestCase {

	return [ ] geoHashTestCase {

		geoHashTestCase {
			precision: 1,
			interval:  Interval(0, 1),
			num:       0.5,
			result:   geohash {
				values: [ ] bool {
					false,
				},
			},
		},
		geoHashTestCase {
			precision: 1,
			interval:  Interval(0, 1),
			num:       0.55,
			result:   geohash {
				values: [ ] bool {
					true,
				},
			},
		},
		geoHashTestCase {
			precision: 5,
			interval:  Interval(0, 1),
			num:       1e-7,
			result:   geohash {
				values: [ ] bool {
					false,
					false,
					false,
					false,
					false,
				},
			},
		},
		geoHashTestCase {
			precision: 2,
			interval:  Interval(0, 1),
			num:       0.7,
			result:   geohash {
				values: [ ] bool {
					true,
					false,
				},
			},
		},
	}

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

	gob := goblin.Goblin(test)

	gob.Describe("Geohash", func ( ) {

		var result geohash

		for _, testCase := range getTestCases( ) {

			result = Geohash(testCase.precision, testCase.interval, testCase.num)

			gob.It("has the expected length", func ( ) {
				gob.Assert(len(result.values)).Equal(len(testCase.result.values))
			})

			gob.It("generates the expected output value", func ( ) {

				gob.Assert(result).Equal(testCase.result)

				fmt.Println("asd")

			})

		}

	})

	gob.Describe("Geohash2d", func ( ) {

		//var result         geohash2d
		var testInterval2d interval2d
		var testPoint      point

		for _, testCase := range getTestCases( ) {

			testInterval2d = interval2d {
				x: testCase.interval,
				y: testCase.interval,
			}

			testPoint = point {
				x: testCase.num,
				y: testCase.num,
			}

			_ = Geohash2d(testCase.precision, testInterval2d, testPoint)

			gob.It("generates the expected output value", func ( ) {


			})

		}

	})

}





func TestGeohashIdempotency (test *testing.T) {

	gob       := goblin.Goblin(test)

	gob.Describe("Geohash", func ( ) {

		var geohash        geohash2d
		var geohash2       geohash2d
		var uintConversion uint64
		var testInterval2d interval2d
		var testPoint      point

		for _, testCase := range getTestCases( ) {

			testInterval2d = interval2d {
				x: testCase.interval,
				y: testCase.interval,
			}

			testPoint = point {
				x: testCase.num,
				y: testCase.num,
			}

			geohash        = Geohash2d(testCase.precision, testInterval2d, testPoint)
			uintConversion = Geohash2dAsUint64(geohash)
			geohash2       = uint64AsGeohash2d(testCase.precision, uintConversion)

			_ = geohash2

		}

	})

}
