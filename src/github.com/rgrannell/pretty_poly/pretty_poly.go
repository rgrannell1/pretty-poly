
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




func writeSolutions (solutionsChannel chan [ ] complex128, args appArguments) {

	for solutions := range solutionsChannel {
		for _, solution := range solutions {

			processSolution(solution, args)

		}
	}

	close(solutionsChannel)

}




func foo (extreme int, order int, processes int) {

	args := appArguments {
		precision:  8,
		dimensions: Interval2d(0, 60000, 0, 6000),
	}

	var waitGroup sync.WaitGroup

	bases := [ ] int { }

	for ith := 0; ith < order; ith++ {
		bases = append(bases, extreme)
	}

	solutions     := make(chan [ ] complex128)
	exploredBases := make([ ] int, order)

	for ith := 0; ith < order; ith++ {
		exploredBases[ith] = (2 * extreme) + 1
	}

	waitGroup.Add(processes)

	for offset := 0; offset < processes; offset++ {

		go func (offset int) {

			defer waitGroup.Done( )

			coefficientCount := productOf(exploredBases)

			for ith := offset; ith < coefficientCount; ith += processes {
				solutions <- solvePolynomial( toCompanionMatrix(toMixedRadix(exploredBases, ith), float64(extreme)) )
			}

		}(offset)

	}

	go writeSolutions(solutions, args)

	waitGroup.Wait( )

}
