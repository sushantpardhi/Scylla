package board

// BitBoard is a 64-bit integer where each bit represents a square on the chessboard.
type BitBoard uint64

// square represents a specific square on the chessboard, numbered from 0 (A1) to 63 (H8).
type square uint

// The constants A1 to H8 represent the squares on the chessboard, starting from the bottom-left (A1) to the top-right (H8).
const (
	A1 square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// BitAt returns a mask with only square s set to 1.
// Shifting 1 left by s moves that single set bit to the square's position.
func BitAt(s square) BitBoard {
	return BitBoard(1) << s
}

// SetBit returns board with square s set.
// OR (|) keeps every bit already set in board and turns on the bit in BitAt(s).
func SetBit(board BitBoard, s square) BitBoard {
	return board | BitAt(s)
}

// ClearBit returns board with square s cleared.
// Bit clear (&^) keeps board's bits except the bit selected by BitAt(s).
func ClearBit(board BitBoard, s square) BitBoard {
	return board &^ BitAt(s)
}

// IsBitSet reports whether square s is set in
// AND (&) retains only bits shared by board and BitAt(s), so a non-zero
// result means the square's bit was set.
func IsBitSet(board BitBoard, s square) bool {
	return board&BitAt(s) != 0
}

// PopCount returns the number of set bits (1s) in the board.
func PopCount(board BitBoard) int {
	count := 0
	for board != 0 {
		board &= board - 1
		count++
	}
	return count
}
