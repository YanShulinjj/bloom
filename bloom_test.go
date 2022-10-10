/* ----------------------------------
*  @author suyame 2022-10-10 19:28:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package bloom

import (
	"fmt"
	"testing"
)

func TestBloom(t *testing.T) {
	bl := NewBloom(256<<22, ExampleFuns)

	for i := 0; i < 10000; i++ {
		bl.Add(fmt.Sprint('a' + i))
	}
	for i := 0; i < 10000; i++ {
		if !bl.Exists(fmt.Sprint('a' + i)) {
			t.Errorf("%s should exits!", fmt.Sprint('a'+i))
		}
	}
	for i := 10001; i < 20000; i++ {
		if bl.Exists(fmt.Sprint('a' + i)) {
			t.Errorf("%s should not exits!", fmt.Sprint('a'+i))
		}
	}
}
