package board

// Position represents the state of a chess game, with separate BitBoards for each piece type and color.
type Position struct {
	WP BitBoard // White Pawn
	WN BitBoard // White Knight
	WB BitBoard // White Bishop
	WR BitBoard // White Rook
	WQ BitBoard // White Queen
	WK BitBoard // White King
	BP BitBoard // Black Pawn
	BN BitBoard // Black Knight
	BB BitBoard // Black Bishop
	BR BitBoard // Black Rook
	BQ BitBoard // Black Queen
	BK BitBoard // Black King
}

// NewGame initializes a Position with the standard starting setup for a chess game.
func NewGame() Position {
	return Position{
		WP: BitAt(A2) | BitAt(B2) | BitAt(C2) | BitAt(D2) | BitAt(E2) | BitAt(F2) | BitAt(G2) | BitAt(H2),
		WN: BitAt(B1) | BitAt(G1),
		WB: BitAt(C1) | BitAt(F1),
		WR: BitAt(A1) | BitAt(H1),
		WQ: BitAt(D1),
		WK: BitAt(E1),
		BP: BitAt(A7) | BitAt(B7) | BitAt(C7) | BitAt(D7) | BitAt(E7) | BitAt(F7) | BitAt(G7) | BitAt(H7),
		BN: BitAt(B8) | BitAt(G8),
		BB: BitAt(C8) | BitAt(F8),
		BR: BitAt(A8) | BitAt(H8),
		BQ: BitAt(D8),
		BK: BitAt(E8),
	}
}

// WhitePieces returns a BitBoard representing all white pieces on the board.
func (position Position) WhitePieces() BitBoard {
	return position.WP | position.WN | position.WB | position.WR | position.WQ | position.WK
}

// BlackPieces returns a BitBoard representing all black pieces on the board.
func (position Position) BlackPieces() BitBoard {
	return position.BP | position.BN | position.BB | position.BR | position.BQ | position.BK
}
