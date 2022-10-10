/* ----------------------------------
*  @author suyame 2022-10-10 19:20:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"bloom"
	"fmt"
)

func main() {
	bl := bloom.NewBloom(1024, bloom.ExampleFuns)
	bl.Add("suyame")
	fmt.Println(bl.Exists("suyame"))
	fmt.Println(bl.Exists("suyame2"))
}
