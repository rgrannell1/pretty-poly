
package pretty_poly




import "math"
import "image"
import "image/color"





type PolynomialImage struct {
	content *image.RGBA
}



/*

*/

func (graph PolynomialImage) ColorModel ( ) color.Model {
	return graph.content.ColorModel( )
}

func (graph PolynomialImage) Bounds ( ) image.Rectangle {
	return graph.content.Bounds( )
}

func (graph PolynomialImage) At(x int, y int) color.Color {
	return graph.content.At(x, y)
}

func (graph PolynomialImage) Set(hash geohash2d, interval interval2d) {

	colour := color.RGBA {
		255,
		0,
		0,
		255,
	}


	solution := hash.AsPoint(interval)

	graph.content.Set(int(solution.x), int(solution.y), colour)

}

func (graph PolynomialImage) SetPoint(point point2d) {

	colour := color.RGBA {
		255,
		0,
		0,
		255,
	}

	graph.content.Set(int(point.x), int(point.y), colour)

}

func (graph PolynomialImage) SetComplex(solution complex128, interval interval2d) {

	realPart := real(solution)
	imagPart := imag(solution)

	colour := color.RGBA {
		255,
		0,
		0,
		255,
	}

	xMin := interval.x.lower
	yMin := interval.y.lower

	xMax := interval.x.upper
	yMax := interval.y.upper

	xDiff := xMax - xMin
	yDiff := yMax - yMin

	xPercent := (realPart - xMin) / xDiff
	yPercent := (imagPart - yMin) / yDiff

	x := math.Floor(xPercent * xMax)
	y := math.Floor(yPercent * yMax)

	graph.content.Set(int(x), int(y), colour)

}
