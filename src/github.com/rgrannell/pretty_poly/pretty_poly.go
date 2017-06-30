
package pretty_poly

import "sync"
import "io"
import "os"
import "bufio"
import "encoding/binary"
import "image"
import "image/png"
import "image/color"





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





func writeManager (filepath string, solutionsChan chan [ ] complex128, writeChan chan uint64) {

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



func SolvePolynomials (extreme int, order int, filepath string) {

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

	go writeManager(filepath, solutionsChan, writeChan)
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




func DrawImage (filepath string) error {

	conn, err := os.Open(filepath)
	defer conn.Close( )

	if err != nil {
		return err
	}

	buffer := make([ ] byte, 8)

	dimensions := interval2d {
		x: interval {
			lower: 0,
			upper: 256,
		},
		y: interval {
			lower: 0,
			upper: 256,
		},
	}

	img := PolynomialImage {
		content: image.NewRGBA(image.Rect(
			int(dimensions.x.lower),
			int(dimensions.y.lower),
			int(dimensions.x.upper),
			int(dimensions.y.upper),
		)),
	}

	for ith := dimensions.x.lower; ith < dimensions.x.upper; ith++ {
		for jth := dimensions.y.lower; jth < dimensions.y.upper; jth++ {

			img.content.Set(int(ith), int(jth), color.RGBA {
				25,
				25,
				25,
				255,
			})

		}
	}

	for {

		count, err := conn.Read(buffer)

		if err != nil && err != io.EOF {
			return err
		}

		if count != 8 {
			break
		}

		solution, err := Uint64AsGeohash2d(8, binary.LittleEndian.Uint64(buffer))

		if err != nil {
			return err
		}

		img.Set(solution, dimensions)

	}

	outConn, outErr := os.Create(filepath + ".png")

	if outErr != nil {
		return outErr
	}

	defer outConn.Close( )

	pngErr := png.Encode(outConn, img)

	if pngErr != nil {
		return pngErr
	}

	outConn.Sync( )

	return nil

}
