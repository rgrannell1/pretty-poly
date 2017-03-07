
package pretty_poly





import "testing"





func BenchmarkSolving1 (bench *testing.B) {
	foo(10, 6, 1)
}

func BenchmarkSolving2 (bench *testing.B) {
	foo(10, 6, 2)
}

func BenchmarkSolving4 (bench *testing.B) {
	foo(10, 6, 4)
}

func BenchmarkSolving8 (bench *testing.B) {
	foo(10, 6, 8)
}

func BenchmarkSolving16 (bench *testing.B) {
	foo(10, 6, 16)
}

func BenchmarkSolving32 (bench *testing.B) {
	foo(10, 6, 32)
}

func BenchmarkSolving64 (bench *testing.B) {
	foo(10, 6, 64)
}
