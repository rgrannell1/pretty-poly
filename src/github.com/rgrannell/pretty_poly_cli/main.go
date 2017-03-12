
package main




import "fmt"
import "github.com/docopt/docopt-go"
//import "github.com/rgrannell/pretty_poly"






func main ( ) {

	usage := `Pretty Poly

	Usage:
		pretty_poly -h | --help
		pretty_poly --version

	Options:
		-h, --help    Show this documentation.
		--version     Show the package version.

	`

	arguments, err := docopt.Parse(usage, nil, true, "Pretty Poly 0.1", false)

	if err != nil {
		panic(err)
	}

	fmt.Println( arguments )

}
