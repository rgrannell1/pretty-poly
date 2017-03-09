
package pretty_poly





import "testing"
import "github.com/franela/goblin"




type geoHashTestCase struct {
	precision int8
	interval  interval
	num       float64
	result    geohash
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
		_ = Geohash2dAsUint64(geohash)
	}

}





func BenchmarkUint64AsGeohash2d (bench *testing.B) {

	interval := Interval2d(0, 60000, 0, 60000)
	point    := point2d {
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
	var geohash2       geohash2d
	var uintConversion uint64
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

	geohash        = Geohash2d(testCase.precision, testInterval2d, testPoint)
	uintConversion = Geohash2dAsUint64(geohash)
	geohash2       = uint64AsGeohash2d(testCase.precision, uintConversion)

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
