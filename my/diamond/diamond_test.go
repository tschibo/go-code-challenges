package diamond

import "fmt"

func ExampleRenderDiamond() {
	fmt.Println(RenderDiamond("A"))
	// Output:
	// A
}

func ExampleRenderDiamond2() {
	fmt.Println(RenderDiamond("D"))
	// Output:
	//    A
	//   B B
	//  C   C
	// D     D
	//  C   C
	//   B B
	//    A
}
