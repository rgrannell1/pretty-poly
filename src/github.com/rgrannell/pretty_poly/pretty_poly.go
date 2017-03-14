
package pretty_poly

import "fmt"
import "sync"
import "io"
import "os"
import "bufio"
import "encoding/binary"



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





func writeManager (filename string, solutionsChan chan [ ] complex128, writeChan chan uint64) {

	conn, err := os.Create("./" + filename)
	defer conn.Close( )

	writer := bufio.NewWriter(conn)
	defer writer.Flush( )

	if err != nil {
		panic(err)
	}

	count := 0

	for solution := range writeChan {

		writer.WriteString(string(solution) + "\n")
		count++

	}

	println(count)

}



func SolvePolynomials (extreme int, order int, filename string) {

	processes := 20

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

	go writeManager(filename, solutionsChan, writeChan)
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




func DrawImage (filename string) {

	conn, err := os.Open("./" + filename)
	defer conn.Close( )

	if err != nil {
		panic(err)
	}

	buffer := make([ ] byte, 8)

	for {

		count, err := conn.Read(buffer)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if count != 8 {
			break
		}

		fmt.Println(
			binary.LittleEndian.Uint64(buffer),
		)

	}

}
