
package pretty_poly

//import "fmt"
import "sync"




func foo (extreme int, order int, processes int) {

	var waitGroup sync.WaitGroup

	bases := [ ] int { }

	for ith := 0; ith < order; ith++ {
		bases = append(bases, extreme)
	}

	solutions := make(chan [ ] complex128)
	outBases  := make([ ] int, order)

	for ith := 0; ith < order; ith++ {
		outBases[ith] = extreme
	}

	waitGroup.Add(processes)

	for offset := 0; offset < processes; offset++ {

		go func (offset int) {

			defer waitGroup.Done( )

			coefficientCount := productOf(outBases)

			for ith := offset; ith < coefficientCount; ith += processes {
				solutions <- solvePolynomial( toCompanionMatrix(toMixedRadix(outBases, ith)) )
			}

		}(offset)

	}

	go func ( ) {

		for _ = range solutions {

		}

	}( )

	waitGroup.Wait( )
	close(solutions)

}
