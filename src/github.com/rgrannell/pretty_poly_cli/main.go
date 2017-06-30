
package main



import "strconv"

import "github.com/docopt/docopt-go"
import "github.com/rgrannell/pretty_poly"




const commandUsage = `
Usage:
	pretty_poly solve --path <path> [--extreme <extreme>] [--order <order>] [--precision <num>]
	pretty_poly draw  --path <path>

	pretty_poly -h | --help
	pretty_poly --version

Description:
	.

Options:
	--path      <path>       asdasd .
	--extreme   <extreme>    The largest integer coefficient to use [default: 5].
	--order     <order>      The order of the polynomial to solve [default: 3].
	--precision <num>        the precision at which to plot [default: 10].
	-h, --help               Show this documentation.
	--version                Show the package version.
`





func solveCliCommand (args map[string] interface { }) error {

	order, parseOrderErr := strconv.ParseInt(args["<order>"].(string), 10, 64)

	if parseOrderErr != nil {
		return parseOrderErr
	}

	extreme, parseIntErr := strconv.ParseInt(args["<extreme>"].(string), 10, 64)

	if parseIntErr != nil {
		return parseIntErr
	}

	runTimeErr := pretty_poly.Solve(
		int(order),
		int(extreme),
		args["<path>"].(string),
	)

	if runTimeErr != nil {
		return runTimeErr
	}

	return nil

}





func drawCliCommand (args map[string] interface { }) error {

	pretty_poly.Draw(args["<path>"].(string))
	return nil

}




func startCommandLine ( ) error {

	args, err := docopt.Parse(commandUsage, nil, true, "Pretty Poly 0.1", false)

	if err != nil {
		return err
	}

	if args["solve"] == true {

		return solveCliCommand(args)

	} else if args["draw"] == true {

		return drawCliCommand(args)

	}

	return nil

}


func main ( ) {

	err := startCommandLine( )

	if err != nil {
		panic(err)
	}

}
