package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ // overflow
	fmt.Println(x)

	// a := int8(255 + 1)
	// _ = a

	var b int8 = 127
	fmt.Println("%d \n", b+1)

	f := float32(math.MaxFloat32)
	fmt.Println(f * 1.2)

	//const i int8 = 300
}
