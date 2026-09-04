package board

import (
	"io"
	"os"
	"testing"
)

// TestBitAt verifies that the BitAt function correctly returns a BitBoard with only the specified square set.
func TestBitAt(t *testing.T) {
	tests := []struct {
		name   string
		square square
		want   BitBoard
	}{
		{"First Square", A1, BitAt(A1)},  // Test the first square (A1) of the board.
		{"Middle Square", D4, BitAt(D4)}, // Test a middle square (D4) of the board.
		{"Last Square", H8, BitAt(H8)},   // Test the last square (H8) of the board.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := BitAt(test.square); got != test.want {
				t.Fatalf("BitAt(%d) = %d, want %d", test.square, got, test.want)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	var board BitBoard

	board = SetBit(board, A1)
	board = SetBit(board, D4)
	board = SetBit(board, H8)

	if !IsBitSet(board, A1) {
		t.Fatal("A1 should be set")
	}
	if !IsBitSet(board, D4) {
		t.Fatal("D4 should be set")
	}
	if !IsBitSet(board, H8) {
		t.Fatal("H8 should be set")
	}

	if PopCount(board) != 3 {
		t.Fatalf("PopCount failed, got %d, want %d", PopCount(board), 3)
	}

	board = SetBit(board, A1)
	if PopCount(board) != 3 {
		t.Fatalf("Setting A1 twice gave PopCount, got %d, want %d", PopCount(board), 3)
	}

}

func TestClearBit(t *testing.T) {
	board := SetBit(BitBoard(0), A1)
	board = SetBit(board, H8)

	board = ClearBit(board, A1)

	if IsBitSet(board, A1) {
		t.Fatal("A1 should be cleared")
	}

	if !IsBitSet(board, H8) {
		t.Fatal("clearing A1 should not clear H8")
	}

	// Clearing an already-clear bit should not change the board.
	before := board
	board = ClearBit(board, A1)
	if board != before {
		t.Fatalf("clearing an unset bit changed board from %d to %d", before, board)
	}
}

func TestIsBitSet(t *testing.T) {
	board := SetBit(BitBoard(0), E4)

	if !IsBitSet(board, E4) {
		t.Fatal("E4 should be set")
	}

	if IsBitSet(board, E5) {
		t.Fatal("E5 should not be set")
	}
}

func TestPopCount(t *testing.T) {
	tests := []struct {
		name  string
		board BitBoard
		want  int
	}{
		{"empty board", 0, 0},
		{"one bit", BitBoard(1) << A1, 1},
		{"three bits", BitBoard(1)<<A1 | BitBoard(1)<<D4 | BitBoard(1)<<H8, 3},
		{"full board", ^BitBoard(0), 64},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := PopCount(test.board); got != test.want {
				t.Fatalf("PopCount(%d) = %d, want %d", test.board, got, test.want)
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	position := NewGame()

	if PopCount(position.WP) != 8 {
		t.Fatalf("white pawns = %d, want 8", PopCount(position.WP))
	}

	if !IsBitSet(position.WK, E1) {
		t.Fatal("white king should be on E1")
	}

	if !IsBitSet(position.BK, E8) {
		t.Fatal("black king should be on E8")
	}

	if PopCount(position.WhitePieces()) != 16 {
		t.Fatalf("white pieces = %d, want 16", PopCount(position.WhitePieces()))
	}

	if PopCount(position.BlackPieces()) != 16 {
		t.Fatalf("black pieces = %d, want 16", PopCount(position.BlackPieces()))
	}
}

func TestPrint(t *testing.T) {
	originalStdout := os.Stdout

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}

	os.Stdout = writer

	Print(BitAt(A1))

	writer.Close()
	os.Stdout = originalStdout

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read Print output: %v", err)
	}
	reader.Close()

	if len(output) == 0 {
		t.Fatal("Print produced no output")
	}
}
