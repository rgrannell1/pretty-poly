
package pretty_poly





import "testing"
import "github.com/franela/goblin"
import "github.com/gonum/matrix/mat64"





func BenchmarkToCompanionMatrix (bench *testing.B) {

	bases := [ ] float64 {100, 100, 100, 100, 100}

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		toCompanionMatrix(bases)
	}

}




func BenchmarkSolvePolynomial (bench *testing.B) {

	bases := [ ] float64 {100, 100, 100, 100, 100}

	bench.StartTimer( )

	for ith := 0; ith < bench.N; ith++ {
		solvePolynomial(toCompanionMatrix(bases))
	}

}





func TestToCompanionMatrix (test *testing.T) {

	var coord0     float64
	var coord1     float64
	var coeff      [ ]float64
	var polyMatrix *mat64.Dense

	gob := goblin.Goblin(test)

	gob.Describe("to companion matrix", func ( ) {

		gob.It("constructs the expected matrices", func ( ) {

			coord0     = float64(+10)
			coord1     = float64(-10)

			coeff      = [ ]float64{ coord0, coord1 }
			polyMatrix = toCompanionMatrix(coeff)

			gob.Assert(polyMatrix.At(0, 1)).Equal(-coord0)
			gob.Assert(polyMatrix.At(1, 1)).Equal(-coord1)

		})

	})

}
