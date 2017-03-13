
	package main




	import "strconv"

	import "github.com/docopt/docopt-go"
	import "github.com/rgrannell/pretty_poly"




	const commandUsage = `
	Usage:
		pretty_poly solve --name <str> [--extreme <extreme>] [--order <order>]
		pretty_poly draw --name <str>

		pretty_poly -h | --help
		pretty_poly --version

	Description:
		.

	Options:
		--name <str>       asdasdasd.
		--extreme <extreme>    The largest integer coefficient to use [default: 5].
		--order <order>        The order of the polynomial to solve [default: 3].
		-h, --help             Show this documentation.
		--version              Show the package version.
	`





	func main ( ) {

		args, err := docopt.Parse(commandUsage, nil, true, "Pretty Poly 0.1", false)

		if err != nil {
			panic(err)
		}

		if args["solve"] == true {

			order, parseOrderErr := strconv.ParseInt(args["--order"].(string), 10, 64)

			if parseOrderErr != nil {
				panic(".")
			}

			extreme, parseIntErr := strconv.ParseInt(args["--extreme"].(string), 10, 64)

			if parseIntErr != nil {
				panic(".")
			}

			pretty_poly.Solve(
				int(order),
				int(extreme),
				args["--name"].(string),
			)

		} else if args["draw"] == true {

			pretty_poly.Draw(args["--name"].(string))

		}

	}
