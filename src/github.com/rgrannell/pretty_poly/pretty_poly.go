
package pretty_poly

import "fmt"
import "sync"
import "encoding/binary"
import "os"



type appArguments struct {

	precision  int8
	dimensions interval2d

}




func processSolution (solution complex128, args appArguments) (uint64, error) {

	argandPoint := point2d {
		x: real(solution),
		y: imag(solution),
	}

	return Geohash2dAsUint64(
		Geohash2d(args.precision, args.dimensions, argandPoint),
	)

}





func writeManager (solutionsChan chan [ ] complex128, writeChan chan uint64) {

	conn, err := os.Create("/tmp/test_0")
	defer conn.Close( )

	if err != nil {
		panic(err)
	}

	for solution := range writeChan {

		encoded := make([ ] byte, 8)
		binary.LittleEndian.PutUint64(encoded, solution)

		conn.Write(encoded)

	}

}



func RunMeRunMe (extreme int, order int, processes int) {

	args := appArguments {
		precision:  8,
		dimensions: Interval2d(-10, 10, -10, 10),
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

	writeChan := make(chan uint64, 100)

	go writeManager(solutionsChan, writeChan)
	go func ( ) {

		for solutions := range solutionsChan {
			for _, solution := range solutions {

				encoded, err := processSolution(solution, args)

				if err != nil {
					panic("aarrgghh!")
				}

				writeChan <- encoded

			}
		}

	}( )

	solveGroup.Wait( )

}

func ReadIn ( ) {

	conn, err := os.Create("/tmp/test_0")
	defer conn.Close( )

	if err != nil {
		panic(err)
	}

	foo := make([ ] byte, 8 * 1000)

	conn.Read(foo)

	fmt.Println( foo )

}



