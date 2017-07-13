
package pretty_poly





import "fmt"
// import "reflect"
import "testing"
import "github.com/franela/goblin"





func TestUtils (test *testing.T) {

	test.Run("toBits", func (test *testing.T) {

		gob := goblin.Goblin(test)

		gob.Describe("toBits", func ( ) {

			gob.It("works for known test cases", func ( ) {

				gob.Assert( toBits(0) ).Equal([ ] bool {false})
				gob.Assert( toBits(1) ).Equal([ ] bool {true})
				gob.Assert( toBits(2) ).Equal([ ] bool {true, false})
				gob.Assert( toBits(3) ).Equal([ ] bool {true, true})
				gob.Assert( toBits(4) ).Equal([ ] bool {true, false, false})
				gob.Assert( toBits(5) ).Equal([ ] bool {true, false, true})

			})

			gob.It("does not panic for random input", func ( ) {

				for num := 0; num < 100; num++ {
					toBits(uint64(num))
				}

			})

		})


	})

}




func TestProductOf (test *testing.T) {

	gob := goblin.Goblin(test)

	gob.Describe("toBits", func ( ) {

		gob.It("product of zero et. al is zero", func ( ) {

			gob.Assert(productOf( [ ]int {0, 0} )).Equal(0)
			gob.Assert(productOf( [ ]int {1, 0, 1} )).Equal(0)

		})

		gob.It("product of one times a number is the number", func ( ) {

			gob.Assert(productOf( [ ]int { } )).Equal(1)
			gob.Assert(productOf( [ ]int {1} )).Equal(1)
			gob.Assert(productOf( [ ]int {1, 5} )).Equal(5)
			gob.Assert(productOf( [ ]int {1, 1, 5} )).Equal(5)

		})

	})

}

func TestFromBitsLittleEndian (test *testing.T) {

	result0 := fromBitsLittleEndian([ ] bool { })
	result1 := fromBitsLittleEndian([ ] bool {true})
	result2 := fromBitsLittleEndian([ ] bool {false, true})
	result3 := fromBitsLittleEndian([ ] bool {true, true})
	result4 := fromBitsLittleEndian([ ] bool {false, false, true})

	if result0 != 0 {
		panic(fmt.Sprintf("mismatched %d, expected %d", result0, 0))
	}

	if result1 != 1 {
		panic(fmt.Sprintf("mismatched %d, expected %d", result1, 1))
	}

	if result2 != 2 {
		panic(fmt.Sprintf("mismatched %d, expected %d", result0, 2))
	}

	if result3 != 3 {
		panic(fmt.Sprintf("mismatched %d, expected %d", result0, 3))
	}

	if result4 != 4 {
		panic(fmt.Sprintf("mismatched %d, expected %d", result0, 4))
	}

}

/*
func TestIntersperseBool (test *testing.T) {

	result0 := IntersperseBool([ ] bool { }, [ ] bool { })
	result1 := IntersperseBool([ ] bool {true}, [ ] bool {false})
	result2 := IntersperseBool([ ] bool {true, true}, [ ] bool {false, false})

	if !reflect.DeepEqual(result0, [ ] bool { }) {
		panic(fmt.Sprintf("mismatched case empty %v", result0))
	}

	if !reflect.DeepEqual(result1, [ ] bool {true, false}) {
		panic(fmt.Sprintf("mismatched case l1 %v", result1))
	}

	if !reflect.DeepEqual(result2, [ ] bool {true, false, true, false}) {
		panic(fmt.Sprintf("mismatched case l2 %v", result2))
	}

	_ = result0
	_ = result1
	_ = result2

}

*/

