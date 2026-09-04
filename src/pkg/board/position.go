package board

type Position struct {
	WP BitBoard
	WN BitBoard
	WB BitBoard
	WR BitBoard
	WQ BitBoard
	WK BitBoard
	BP BitBoard
	BN BitBoard
	BB BitBoard
	BR BitBoard
	BQ BitBoard
	BK BitBoard
}

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

func (position Position) WhitePieces() BitBoard {
	return position.WP | position.WN | position.WB | position.WR | position.WQ | position.WK
}

func (position Position) BlackPieces() BitBoard {
	return position.BP | position.BN | position.BB | position.BR | position.BQ | position.BK
}
