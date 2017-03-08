
package pretty_poly





import "testing"





func BenchmarkSolving1 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 1)
	}

}

func BenchmarkSolving2 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 2)
	}

}

func BenchmarkSolving4 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 4)
	}

}

func BenchmarkSolving8 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 8)
	}

}

func BenchmarkSolving16 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 16)
	}

}

func BenchmarkSolving32 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 32)
	}

}

func BenchmarkSolving64 (bench *testing.B) {

	for ith := 0; ith < bench.N; ith++ {
		foo(5, 5, 64)
	}

}
