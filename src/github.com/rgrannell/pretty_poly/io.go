
package pretty_poly

import "io"
import "os"
import "fmt"
import "bufio"
import "encoding/binary"
import "github.com/alash3al/goemitter"





func writeGeocodeManager (filepath string, solutionsChan chan [ ] complex128, writeChan chan uint64) {

	conn, err := os.Create(filepath)
	defer conn.Close( )

	writer := bufio.NewWriter(conn)
	defer writer.Flush( )

	if err != nil {
		panic(err)
	}

	count := 0

	for solution := range writeChan {

		buffer := make([ ] byte, 8, 8)
		binary.LittleEndian.PutUint64(buffer, solution)

		for _, uintByte := range buffer {
			writer.WriteByte(uintByte)
		}

		count++

	}

}





func emitGeocodeWrites (solutionsChan chan [ ] complex128, precision int8, writeChan chan uint64) {

	dimensions := Interval2d(-10, 10, -10, 10)

	for solutions := range solutionsChan {
		for _, solution := range solutions {


			argandPoint := point2d {
				x: real(solution),
				y: imag(solution),
			}

			encoded, err := Geohash2dAsUint64(
				Geohash2d(precision, dimensions, argandPoint),
			)

			if err != nil {
				panic("aarrgghh!")
			}

			writeChan <- encoded

		}
	}

}




/*
	given a
*/

func writeGeocodeSolutions (filepath string, solutionsChan chan [ ] complex128, precision int8, logger *Emitter.Emitter) {

	writeChan := make(chan uint64, 100)

	go writeGeocodeManager(filepath, solutionsChan, writeChan)
	go emitGeocodeWrites(solutionsChan, precision, writeChan)

}





func writeCartesianManager (filepath string, solutionsChan chan [ ] complex128, writeChan chan complex128) {

	conn, err := os.Create(filepath)
	defer conn.Close( )

	writer := bufio.NewWriter(conn)
	defer writer.Flush( )

	if err != nil {
		panic(err)
	}

	count := 0

	for solution := range writeChan {

		writer.WriteString(fmt.Sprintf("%f+%fi\n", real(solution), imag(solution)))
		count++

	}

}





func emitCartesianWrites (solutionsChan chan [ ] complex128, precision int8, writeChan chan complex128) {

	for solutions := range solutionsChan {
		for _, solution := range solutions {
			writeChan <- solution
		}
	}

}





func writeCartesianSolutions (filepath string, solutionsChan chan [ ] complex128, precision int8, logger *Emitter.Emitter) {

	writeChan := make(chan complex128, 100)

	go writeCartesianManager(filepath, solutionsChan, writeChan)
	go emitCartesianWrites(solutionsChan, precision, writeChan)

}





func readCartesianSolutions (solutionConn *os.File) (chan error, chan complex128) {

	solutions := make(chan complex128, 1)
	errs      := make(chan error, 1)

	go func ( ) {

		for {

			count, err := solutionConn.Read(buffer)

			if err != nil && err != io.EOF {
				errs <- err
			}

			if count != 8 {
				break
			}

			solution, err := Uint64AsGeohash2d(8, binary.LittleEndian.Uint64(buffer))

			if err != nil {
				errs <- err
			} else {
				solutions <- solution
			}

		}

		close(solutions)
		close(errs)

	}( )

	return errs, solutions

}





/*
	given a file pointer, return an error channel and a
	channel of 2d geocode-encoded solutions.
*/

func readGeocodeSolutions (solutionConn *os.File) (chan error, chan geohash2d) {

	buffer    := make([ ] byte, 8)
	solutions := make(chan geohash2d, 1)
	errs      := make(chan error, 1)

	go func ( ) {

		for {

			count, err := solutionConn.Read(buffer)

			if err != nil && err != io.EOF {
				errs <- err
			}

			if count != 8 {
				break
			}

			solution, err := Uint64AsGeohash2d(8, binary.LittleEndian.Uint64(buffer))

			if err != nil {
				errs <- err
			} else {
				solutions <- solution
			}

		}

		close(solutions)
		close(errs)

	}( )

	return errs, solutions

}
