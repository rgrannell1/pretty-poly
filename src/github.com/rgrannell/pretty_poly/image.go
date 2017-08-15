
package pretty_poly




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
