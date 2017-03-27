
package pretty_poly





import "testing"
import "math"
import "math/rand"
import "github.com/franela/goblin"
import "fmt"




type geoHashTestCase struct {
	precision int8
	interval  interval
	num       float64
	result    geohash
}





func geohashCreationTestCase ( ) (float64, [ ] bool, interval) {

	precision := rand.Intn(8) + 1
	extreme   := 10 * float64(rand.Intn(10) + 1)
	interval  := Interval(0, +extreme)

	num   := (interval.upper - interval.lower) / 2
	bools := make([ ] bool, precision, precision)

	for ith := 0; ith < precision; ith++ {

		division := (interval.upper - interval.lower) / math.Pow(2, float64(ith + 2))

		splitUpwards := rand.Intn(2) == 1

		bools[ith] = splitUpwards

		if splitUpwards {
			num += division
		} else {
			num -= division
		}

	}

	return num, bools, interval

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
	point    := point2d {
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
	point    := point2d {
		x: 30000,
		y: 30000,
	}

	geohash := Geohash2d(8, interval, point)

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		_, _ = Geohash2dAsUint64(geohash)
	}

}





func BenchmarkUint64AsGeohash2d (bench *testing.B) {

	interval := Interval2d(0, 60000, 0, 60000)
	point    := point2d {
		x: 30000,
		y: 30000,
	}

	geohash    := Geohash2d(8, interval, point)
	hashInt, _ := Geohash2dAsUint64(geohash)

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		Uint64AsGeohash2d(8, hashInt)
	}

}




func runGeohashEqualityGeohashTest (gob *goblin.G, testCase geoHashTestCase) {

	result := Geohash(testCase.precision, testCase.interval, testCase.num)

	gob.It("has the expected length", func ( ) {
		gob.Assert(len(result.values)).Equal(len(testCase.result.values))
	})

	gob.It("generates the expected output value", func ( ) {
		gob.Assert(result).Equal(testCase.result)
	})

}

func runGeohash2dEqualityGeohashTest (gob *goblin.G, testCase geoHashTestCase) {

	testInterval2d := interval2d {
		x: testCase.interval,
		y: testCase.interval,
	}

	testPoint := point2d {
		x: testCase.num,
		y: testCase.num,
	}

	result := Geohash2d(testCase.precision, testInterval2d, testPoint)

	gob.It("has the expected x length", func ( ) {
		gob.Assert(len(result.xs)).Equal(len(testCase.result.values))
	})

	gob.It("has the expected y length", func ( ) {
		gob.Assert(len(result.ys)).Equal(len(testCase.result.values))
	})

	square := testCase.result.AddYAxis(testCase.result)

	gob.It("generates the expected output value", func ( ) {
		gob.Assert(result).Equal(square)
	})


}





func TestGeohashCreation (test *testing.T) {

	gob := goblin.Goblin(test)

	testCase0 := geoHashTestCase {
		precision: 1,
		interval:  Interval(0, 1),
		num:       0.5,
		result:   geohash {
			values: [ ] bool {
				false,
			},
		},
	}

	testCase1 := geoHashTestCase {
		precision: 1,
		interval:  Interval(0, 1),
		num:       0.55,
		result:   geohash {
			values: [ ] bool {
				true,
			},
		},
	}

	testCase2 := geoHashTestCase {
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
	}

	testCase3 := geoHashTestCase {
		precision: 2,
		interval:  Interval(0, 1),
		num:       0.7,
		result:   geohash {
			values: [ ] bool {
				true,
				false,
			},
		},
	}

	gob.Describe("Random known geohash's match", func ( ) {

		for ith := 0; ith < 1000; ith++ {

			num, expectedGeohash, interval := geohashCreationTestCase( )
			precision     := int8(len(expectedGeohash))
			actualGeohash := Geohash(precision, interval, num).values

			if len(expectedGeohash) != len(actualGeohash) {

				fmt.Println(expectedGeohash)
				fmt.Println(actualGeohash)

				panic("expected and actual geohash mismatched.")

			}

			for jth := 0; jth < len(expectedGeohash); jth++ {

				if actualGeohash[jth] != expectedGeohash[jth] {

					fmt.Println(num)
					fmt.Println(interval)
					fmt.Println(expectedGeohash)
					fmt.Println(actualGeohash)

					panic("expected and actual geohash mismatched.")

				}

			}

		}

	})

	gob.Describe("Geohash", func ( ) {

		runGeohashEqualityGeohashTest(gob, testCase0)
		runGeohashEqualityGeohashTest(gob, testCase1)
		runGeohashEqualityGeohashTest(gob, testCase2)
		runGeohashEqualityGeohashTest(gob, testCase3)

	})

	gob.Describe("Geohash2d", func ( ) {

		runGeohash2dEqualityGeohashTest(gob, testCase0)
		runGeohash2dEqualityGeohashTest(gob, testCase1)
		runGeohash2dEqualityGeohashTest(gob, testCase2)
		runGeohash2dEqualityGeohashTest(gob, testCase3)

	})

}





func runIdempotency (gob *goblin.G, testCase geoHashTestCase) {

	var geohash        geohash2d
	var testInterval2d interval2d
	var testPoint      point2d

	testInterval2d = interval2d {
		x: testCase.interval,
		y: testCase.interval,
	}

	testPoint = point2d {
		x: testCase.num,
		y: testCase.num,
	}

	geohash            = Geohash2d(testCase.precision, testInterval2d, testPoint)
	uintConversion, _ := Geohash2dAsUint64(geohash)
	geohash2, geoErr  := Uint64AsGeohash2d(testCase.precision, uintConversion)

	if geoErr != nil {
		panic(geoErr)
	}

	gob.Assert(geohash).Equal(geohash2)

}

func TestGeohashIdempotency (test *testing.T) {

	gob := goblin.Goblin(test)



	testCase0 := geoHashTestCase {
		precision: 1,
		interval:  Interval(0, 1),
		num:       0.5,
		result:   geohash {
			values: [ ] bool {
				false,
			},
		},
	}

	testCase1 := geoHashTestCase {
		precision: 1,
		interval:  Interval(0, 1),
		num:       0.55,
		result:   geohash {
			values: [ ] bool {
				true,
			},
		},
	}

	testCase2 := geoHashTestCase {
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
	}

	testCase3 := geoHashTestCase {
		precision: 2,
		interval:  Interval(0, 1),
		num:       0.7,
		result:   geohash {
			values: [ ] bool {
				true,
				false,
			},
		},
	}

	gob.Describe("Geohash", func ( ) {

		gob.It("test that uint <-> geohash conversion is idempotent", func ( ) {

			runIdempotency(gob, testCase0)
			runIdempotency(gob, testCase1)
			runIdempotency(gob, testCase2)
			runIdempotency(gob, testCase3)

		})

	})

}
