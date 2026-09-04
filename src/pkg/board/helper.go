package board

import "fmt"

func Print(board BitBoard) {
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)

		for file := range 8 {
			s := square(rank*8 + file)

			if IsBitSet(board, s) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println("  A B C D E F G H")
}
