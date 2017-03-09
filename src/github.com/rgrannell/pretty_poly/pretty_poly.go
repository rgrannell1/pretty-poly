
package pretty_poly

//import "fmt"
import "sync"



type appArguments struct {

	precision  int8
	dimensions interval2d

}




func processSolution (solution complex128, args appArguments) geohash2d {

	argandPoint := point2d {
		x: real(solution),
		y: imag(solution),
	}

	return Geohash2d(args.precision, args.dimensions, argandPoint)

}





func writeManager (solutionsChan chan [] complex128, solveGroup sync.WaitGroup, args appArguments, processes int) {


}



func foo (extreme int, order int, processes int) {

	args := appArguments {
		precision:  8,
		dimensions: Interval2d(0, 60000, 0, 6000),
	}

	var solveGroup sync.WaitGroup

	bases := [ ] int { }

	for ith := 0; ith < order; ith++ {
		bases = append(bases, extreme)
	}

	solutionsChan := make(chan [ ] complex128, 100)
	exploredBases := make([ ] int, order)

	for ith := 0; ith < order; ith++ {
		exploredBases[ith] = (2 * extreme) + 1
	}

	/*
	.
	*/

	for offset := 0; offset < processes; offset++ {

		solveGroup.Add(1)

		go func (offset int) {


			defer solveGroup.Done( )

			coefficientCount := productOf(exploredBases)

			for ith := offset; ith < coefficientCount; ith += processes {
				solutionsChan <- solvePolynomial( toCompanionMatrix(toMixedRadix(exploredBases, ith), float64(extreme)) )
			}

		}(offset)

	}

	go func ( ) {

		for solutions := range solutionsChan {
			for _, solution := range solutions {

				_ = processSolution(solution, args)

			}
		}

	}( )

	solveGroup.Wait( )

}
