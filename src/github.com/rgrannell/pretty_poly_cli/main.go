
package main



import "fmt"
import "strconv"

import "github.com/docopt/docopt-go"
import "github.com/rgrannell/pretty_poly"




const commandUsage = `
Usage:
	pretty_poly solve --path <path> [--extreme <extreme>] [--order <order>] [--precision <precision>]
	pretty_poly draw  --path <path> [--precision <precision>]

	pretty_poly -h | --help
	pretty_poly --version

Description:
	.

Options:
	--path      <path>       asdasd .
	--extreme   <extreme>    The largest integer coefficient to use [default: 5].
	--order     <order>      The order of the polynomial to solve   [default: 3].
	--precision <precision>  the precision at which to plot         [default: 10].
	-h, --help               Show this documentation.
	--version                Show the package version.
`





func solveCliCommand (args map[string] interface { }) error {

	precision, parsePrecisionErr := strconv.ParseInt(args["<precision>"].(string), 10, 8)

	if parsePrecisionErr != nil {
		return parsePrecisionErr
	}

	order, parseOrderErr := strconv.ParseInt(args["<order>"].(string), 10, 64)

	if parseOrderErr != nil {
		return parseOrderErr
	}

	extreme, parseIntErr := strconv.ParseInt(args["<extreme>"].(string), 10, 64)

	if parseIntErr != nil {
		return parseIntErr
	}

	path := args["<path>"].(string)

	runTimeErr := pretty_poly.Solve(int(order), int(extreme), precision, path)

	if runTimeErr != nil {
		return runTimeErr
	}

	return nil

}





func drawCliCommand (args map[string] interface { }) error {

	precision, parsePrecisionErr := strconv.ParseInt(args["<precision>"].(string), 10, 8)

	if parsePrecisionErr != nil {
		return parsePrecisionErr
	}


	return pretty_poly.Draw(args["<path>"].(string), precision)


}




func startCommandLine ( ) error {

	args, err := docopt.Parse(commandUsage, nil, true, "Pretty Poly 0.1", false)

	if err != nil {
		return err
	}

	for key, value := range args {
 	   fmt.Println("Key:", key, "Value:", value)
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
