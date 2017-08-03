
package pretty_poly

import "log"
import "sync"
import "math"
import "io"
import "os"
import "encoding/binary"
import "image"
import "image/png"
import "image/color"
import "github.com/alash3al/goemitter"



func solveRange (offset int, exploredBases [ ] int, extreme int, processes int) chan [ ]complex128 {

	coefficientCount := productOf(exploredBases)
	solutions        := make(chan [ ]complex128, processes)

	go func ( 	) {

		for ith := offset; ith < coefficientCount; ith += processes {
			solutions <- solvePolynomial( toCompanionMatrix(toMixedRadix(exploredBases, ith), float64(extreme)) )
		}

	}( )

	return solutions

}





func mergeSolutions (inputs [ ] chan [ ]complex128) chan [ ]complex128 {

	var mergeGroup sync.WaitGroup

	output := make(chan [ ]complex128)

	mergeGroup.Add(len(inputs))

	for _, input := range inputs {

		go func (input chan [ ]complex128) {

			for solution := range input {
				output <- solution
			}

			mergeGroup.Done( )

		}(input)

	}

	go func ( ) {

		mergeGroup.Wait( )
		close(output)

	}( )

	return output

}





func SolvePolynomials (extreme int, order int, filepath string, precision int8) {

	processes := 20

//	args := appArguments {
//		precision:  precision,
//		dimensions: Interval2d(-10, 10, -10, 10),
//	}

	bases := [ ] int { }

	for ith := 0; ith < order; ith++ {
		bases = append(bases, extreme)
	}

	exploredBases := make([ ] int, order)

	for ith := 0; ith < order; ith++ {
		exploredBases[ith] = (2 * extreme) + 1
	}

	solutions := make([ ] chan [ ]complex128, processes)

	func ( ) {

		for offset := 0; offset < processes; offset ++ {
			solutions = append(solutions, solveRange(offset, exploredBases, extreme, processes))
		}

	}( )


	for x := range mergeSolutions(solutions) {
		println(x)
	}

//	writeSolutions(filepath, solutions, args)

}




func DrawImage (solutionPath string, precision float64) error {

	conn, err := os.OpenFile("/var/log/pretty-poly.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)

	if err != nil {
		log.SetOutput(conn)
	}

	defer conn.Close( )

	buffer := make([ ] byte, 8)
	logger := Emitter.Construct( )


	logger.On(EVENT_DRAW_IMAGE, func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "started drawing image.",
		})

	})

	logger.On(EVENT_DRAW_READ, func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "reading lines from file.",
		})

	})

	logger.On(EVENT_DRAW_READ_DONE, func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "finished reading lines between files.",
		})

	})

	logger.On(EVENT_DRAWN, func (arg ...interface{ }) {

		log.Println(Log {
			level: "info",
			user_message: "drawing done.",
		})

	})





	logger.EmitSync(EVENT_DRAW_IMAGE)
	logger.EmitSync(EVENT_DRAW_READ)
	logger.EmitSync(EVENT_DRAW_READ_DONE)
	logger.EmitSync(EVENT_DRAWN)

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

	logger.EmitSync(EVENT_DRAW_IMAGE)

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

	logger.EmitSync(EVENT_DRAW_READ_DONE)

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

	logger.EmitSync(EVENT_DRAWN)

	return nil

}
