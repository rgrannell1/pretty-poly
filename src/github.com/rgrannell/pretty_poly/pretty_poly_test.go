
package pretty_poly




import "os"
import "io/ioutil"
import "testing"
import "errors"
import "github.com/franela/goblin"





func TestPrettyPoly (test *testing.T) {

	gob          := goblin.Goblin(test)
	tmpFile, err := ioutil.TempFile("/tmp/", "pretty_poly")
	tmpFileName  := tmpFile.Name( )

	if err != nil {
		panic(err)
	}

	gob.Describe("pretty_poly.SolvePolynomials", func ( ) {

		gob.It("creates an output file.", func ( ) {

			SolvePolynomials(5, 5, tmpFileName)

			if _, err := os.Stat(tmpFileName); os.IsNotExist(err) {
				panic(errors.New("file " + tmpFileName + " does not exist."))
			}

		})

	})

	gob.Describe("pretty_poly.DrawImage", func ( ) {

		gob.It("creates an output image.", func ( ) {

			DrawImage(tmpFileName)

			if _, err := os.Stat(tmpFileName); os.IsNotExist(err) {
				panic(errors.New("file " + tmpFileName + " does not exist."))
			}

		})

	})

}
