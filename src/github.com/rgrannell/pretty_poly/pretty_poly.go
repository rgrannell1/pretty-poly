
package pretty_poly

import "log"
import "sync"
import "math"
import "io"
import "os"
import "bufio"
import "encoding/binary"
import "image"
import "image/png"
import "image/color"
import "github.com/alash3al/goemitter"





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



func SolvePolynomials (extreme int, order int, filepath string, precision int8) {

	processes := 20

	args := appArguments {
		precision:  precision,
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




func DrawImage (solutionPath string, precision float64) error {

	conn, err := os.OpenFile("/var/log/pretty-poly.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer conn.Close( )
	log.SetOutput(conn)

	buffer := make([ ] byte, 8)
	logger := Emitter.Construct( )

	logger.On("draw-image", func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "started drawing image.",
		})

	})

	logger.EmitSync("draw-image")

	dimensions := interval2d {
		x: interval {
			lower: 0,
			upper: math.Pow(2, precision),
		},
		y: interval {
			lower: 0,
			upper: math.Pow(2, precision),
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

	solutionConn, err := os.Open(solutionPath)
	defer solutionConn.Close( )

	if err != nil {
		return err
	}

	for {

		count, err := solutionConn.Read(buffer)

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

	outConn, outErr := os.Create(solutionPath + ".png")

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
