
package pretty_poly

import "fmt"
import "log"
import "sync"
import "math"
import "time"
import "io"
import "os"
import "bufio"
import "encoding/binary"
import "image"
import "image/png"
import "image/color"
import "github.com/alash3al/goemitter"




var logger *Emitter.Emitter




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





func emitWrites (solutionsChan chan [ ] complex128, precision int8, writeChan chan uint64) {

	args := appArguments {
		precision:  precision,
		dimensions: Interval2d(-10, 10, -10, 10),
	}

	for solutions := range solutionsChan {
		for _, solution := range solutions {

			encoded, err := processSolution(solution, args)

			if err != nil {
				panic("aarrgghh!")
			}

			writeChan <- encoded

		}
	}

}





func writeGeocodeSolutions (filepath string, solutionsChan chan [ ] complex128, precision int8, logger *Emitter.Emitter) {

	writeChan := make(chan uint64, 100)

	go writeManager(filepath, solutionsChan, writeChan)
	go emitWrites(solutionsChan, precision, writeChan)

}

func startSolutionWorkers (extreme int, order int, precision int8, logger *Emitter.Emitter) (chan [ ] complex128, *sync.WaitGroup) {

	processes := 20
	solved    := 0
	startTime := time.Now( )

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

		log.Println(Log {
			level: "debug",
			user_message: fmt.Sprintf("solution %v started.", offset),
		})

		solveGroup.Add(1)

		go func (offset int) {

			defer solveGroup.Done( )
			defer log.Println(Log {
				level: "debug",
				user_message: fmt.Sprintf("solution %v finshed.", offset),
			})

			coefficientCount := productOf(exploredBases)

			for ith := offset; ith < coefficientCount; ith += processes {

				solved++
				solutionsChan <- solvePolynomial( toCompanionMatrix(toMixedRadix(exploredBases, ith), float64(extreme)) )

				if (solved % 1e5 == 0) {

					elapsedTime := time.Since(startTime).Seconds( )
					rate        := float64(solved) / elapsedTime

					log.Println(Log {
						level: "info",
						user_message: fmt.Sprintf("%v solved (%v per second)", solved, int(rate)),
					})

				}

			}

		}(offset)

	}

	return solutionsChan, &solveGroup
}






func SolvePolynomials (extreme int, order int, filepath string, precision int8) error {

	conn, err := os.OpenFile("./baz.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer conn.Close( )

	multiWriter := io.MultiWriter(os.Stdout, conn)

	log.SetOutput(multiWriter)

	logger = Emitter.Construct( )

	solutionsChan, solveGroup := startSolutionWorkers(extreme, order, precision, logger)
	writeGeocodeSolutions(filepath, solutionsChan, precision, logger)

	solveGroup.Wait( )

	return nil

}

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




// Display
func DrawImage (solutionPath string, precision float64) error {

	conn, err := os.OpenFile("/var/log/pretty-poly.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer conn.Close( )

	multiWriter := io.MultiWriter(os.Stdout, conn)

	log.SetOutput(multiWriter)

	logger = Emitter.Construct( )

	logger.On("EVENT_DRAW_IMAGE", func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "started drawing image.",
		})

	})

	logger.On("EVENT_DRAW_READ", func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "reading lines from file.",
		})

	})

	logger.On("EVENT_DRAW_READ_DONE", func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "finished reading lines between files.",
		})

	})

	logger.On("EVENT_DRAWN", func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "drawing done.",
		})

	})





	logger.EmitSync("EVENT_DRAW_IMAGE")
	logger.EmitSync("EVENT_DRAW_READ")
	logger.EmitSync("EVENT_DRAW_READ_DONE")
	logger.EmitSync("EVENT_DRAWN")

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

	logger.EmitSync("EVENT_DRAW_IMAGE")

	errs, solutions := readGeocodeSolutions(solutionConn)

	for {

		select {
		case err := <-errs:
			return err

		case solution, ok := <- solutions:

			if !ok {
				solutions = nil
			}

			img.Set(solution, dimensions)
		}

		if solutions == nil {
			break
		}

	}

	logger.EmitSync("EVENT_DRAW_READ_DONE")

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

	logger.EmitSync("EVENT_DRAWN")

	return nil

}
