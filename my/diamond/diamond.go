package diamond

import "fmt"

func RenderDiamond(letter string) string {
	// Result string
	var res string

	// Starting rune
	start := 'A'
	starti := int(start)

	// Ending rune
	end := letter[0]
	endi := int(end)

	// Margings & paddings
	left := endi - starti
	mid := 0

	// Upper part of diamond
	for l := starti; l <= endi; l++ {
		for n := 0; n < left; n++ {
			//fmt.Printf(" ")
			res += " "
		}
		//fmt.Printf("%c", rune(l))
		res += fmt.Sprintf("%c", rune(l))
		if mid > 0 {
			for n := 0; n < mid; n++ {
				//fmt.Printf(" ")
				res += " "
			}
			//fmt.Printf("%c\n", rune(l))
			res += fmt.Sprintf("%c\n", rune(l))
			mid += 2
		} else {
			//fmt.Printf("\n")
			res += "\n"
			mid += 1
		}
		left -= 1
	}
	left = 1
	mid -= 4
	// Lower part of diamond
	for l := endi - 1; l >= starti; l-- {
		for n := 0; n < left; n++ {
			//fmt.Printf(" ")
			res += " "
		}
		//fmt.Printf("%c", rune(l))
		res += fmt.Sprintf("%c", rune(l))
		if mid > 0 {
			for n := 0; n < mid; n++ {
				fmt.Printf(" ")
				res += " "
			}
			//fmt.Printf("%c\n", rune(l))
			res += fmt.Sprintf("%c\n", rune(l))
			mid -= 2
		} else {
			//fmt.Printf("\n")
			res += "\n"
			mid -= 1
		}
		left += 1
	}

	return res
}
