
package pretty_poly




import "os"
import "io/ioutil"
import "testing"
import "errors"
import "github.com/franela/goblin"
import "image"
import "math"





func TestPrettyPoly (test *testing.T) {

	gob          := goblin.Goblin(test)
	tmpFile, err := ioutil.TempFile("/tmp/", "pretty_poly")
	tmpFileName  := tmpFile.Name( )
	precision    := 8.0

	if err != nil {
		panic(err)
	}

	gob.Describe("pretty_poly.SolvePolynomials", func ( ) {

		SolvePolynomials(5, 5, tmpFileName)

		gob.It("creates an output file.", func ( ) {

			if _, err := os.Stat(tmpFileName); os.IsNotExist(err) {
				panic(errors.New("file " + tmpFileName + " does not exist."))
			}

		})

		gob.It("creates an non-empty output file.", func ( ) {

			info, err := os.Stat(tmpFileName);

			if err != nil {
				panic(err)
			}

			size := info.Size( )

			if size == 0 {
				panic(errors.New("empty file."))
			}

		})

	})

	gob.Describe("pretty_poly.DrawImage", func ( ) {

		gob.It("creates an output image.", func ( ) {

			DrawImage(tmpFileName, precision)

			if _, err := os.Stat(tmpFileName ); os.IsNotExist(err) {
				panic(errors.New("file " + tmpFileName + " does not exist."))
			}

		})

		gob.It("creates an non-empty output file.", func ( ) {

			info, err := os.Stat(tmpFileName + ".png");

			if err != nil {
				panic(err)
			}

			size := info.Size( )

			if size == 0 {
				panic(errors.New("empty file."))
			}

		})

		gob.It("creates a file with the correct dimension.", func ( ) {

			file, err := os.Open(tmpFileName + ".png");

			defer file.Close( )

			if err != nil {
				panic(err)
			}

			image, _, err := image.DecodeConfig(file)

			if err != nil {
				panic(err)
			}

			expectedDimension := math.Pow(float64(2), float64(precision))

			if float64(image.Width) != expectedDimension {
				panic(errors.New("invalid image width " + string(image.Width)))
			}

			if float64(image.Height) != expectedDimension {
				panic(errors.New("invalid image height " + string(image.Height)))
			}

		})

	})

}
