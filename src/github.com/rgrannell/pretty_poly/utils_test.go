
package pretty_poly





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
