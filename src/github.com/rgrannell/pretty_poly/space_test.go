
package pretty_poly





import "testing"
import "github.com/franela/goblin"




type toMixedRadixTestCase struct {
	bases [ ]int
	num   int
}





func BenchmarkToMixedRadix (bench *testing.B) {

	bases := [ ] int {100, 100, 100, 100, 100}

	for ith := 0; ith < bench.N; ith++ {
		toMixedRadix(bases, 100)
	}

}






func TestToMixedRadix (test *testing.T) {

	var output [ ]float64
	gob    := goblin.Goblin(test)

	gob.Describe("toMixedRadix", func ( ) {

		gob.It("enumerates a space", func ( ) {

			for ith := 0; ith < 100; ith++ {

				output = toMixedRadix([]int {10, 10}, ith)

				switch ith {
					case 0:
						gob.Assert(output[0]).Equal(float64(0))
						gob.Assert(output[1]).Equal(float64(0))
					case 1:
						gob.Assert(output[0]).Equal(float64(0))
						gob.Assert(output[1]).Equal(float64(1))
					case 99:
						gob.Assert(output[0]).Equal(float64(9))
						gob.Assert(output[1]).Equal(float64(9))
				}

			}

		})

	})

}
