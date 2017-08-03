
package pretty_poly





import "os"
//import "sync"
import "bufio"
import "encoding/binary"





type appArguments struct {
	precision  int8
	dimensions interval2d
}






func writeGeoCodeHandler (filepath string, solutionsChan chan [ ] complex128, writeChan chan uint64) {

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





func encodeSolutionAsGeoHash (solution complex128, args appArguments) (uint64, error) {

	argandPoint := point2d {
		x: real(solution),
		y: imag(solution),
	}

	return Geohash2dAsUint64(
		Geohash2d(args.precision, args.dimensions, argandPoint),
	)

}

func writeSolutions (filepath string, solutionsChan chan[ ] complex128, args appArguments) error {

	writer := make(chan uint64)
	errs   := make(chan error, 1)

//	go func ( ) {

		for solutions := range solutionsChan {
			for _, solution := range solutions {

				encoded, err := encodeSolutionAsGeoHash(solution, args)

				if err != nil {
					errs   <- err
				} else {
					writer <- encoded
				}

			}
		}

		close(writer)
		close(errs)

//	}( )

//	go writeGeoCodeHandler(filepath, solutionsChan, writer)

	for x := range writer {
		if (x != x) {
			print(',')
		}
	}

	return nil

}
