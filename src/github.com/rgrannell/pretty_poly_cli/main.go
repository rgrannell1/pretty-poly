
package main




import "github.com/docopt/docopt-go"
import "github.com/rgrannell/pretty_poly"





const commandUsage = `
Usage:
	pretty_poly solve --name <str> [--extreme <num>] [--order <num>]

	pretty_poly -h | --help
	pretty_poly --version

Description:
	.

Options:
	--name <str>       .

	--extreme <num>    The largest integer coefficient to use [default: 5].
	--order <num>      The order of the polynomial to solve [default: 3].

	-h, --help         Show this documentation.
	--version          Show the package version.
`





func main ( ) {

	args, err := docopt.Parse(commandUsage, nil, true, "Pretty Poly 0.1", false)

	if err != nil {
		panic(err)
	}

	if args["solve"] == true {

		pretty_poly.Solve(pretty_poly.SolveArgs {
			Order:    args["--order"].(int),
			Extreme:  args["--extreme"].(int),
			FileName: args["--name"].(string),
		})

	} else if args["draw"] == true {

		pretty_poly.Draw(pretty_poly.DrawArgs {

		})

	}

}
