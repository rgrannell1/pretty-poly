
package pretty_poly





import "fmt"
import "testing"
import "github.com/franela/goblin"





func runToBitsTest (gob *goblin.G, num uint64, expected [ ] bool ) {

	gob.It("converts to binary correctly", func ( ) {
		gob.Assert(toBits(num)).Equal(expected)
	})

}

func TestToBits (test *testing.T) {

	gob := goblin.Goblin(test)

	gob.Describe("toBits", func ( ) {

		runToBitsTest(gob, 1, [ ] bool {true} )
		runToBitsTest(gob, 2, [ ] bool {true, false} )
		runToBitsTest(gob, 3, [ ] bool {true, true} )
		runToBitsTest(gob, 4, [ ] bool {true, false, false} )
		runToBitsTest(gob, 5, [ ] bool {true, false, true} )
		runToBitsTest(gob, 6, [ ] bool {true, true, false} )
		runToBitsTest(gob, 7, [ ] bool {true, true, true} )

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
