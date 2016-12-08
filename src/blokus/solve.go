package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/michaelbpaulson/blokus"
)

type usage struct {
	piece *blokus.Piece
	row   int
	col   int
}

func (u *usage) String() string {
	return fmt.Sprintf("Row(%v) Col(%v)\n%v", u.row, u.col, u.piece)
}

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	searchType := "slow"

	if len(os.Args) > 2 {
		searchType = os.Args[2]
	}

	board := blokus.NewBoard(size)
	pieces := blokus.GetPieces()

	if searchType == "fast" {
		solveFAI(board, pieces, make([]*usage, len(pieces)), 0, 0)
	} else {
		solveSC(board, pieces, make([]*usage, len(pieces)), 0, 0)
	}
}

func quickClone(remaining []*blokus.PieceGroup, remove int) []*blokus.PieceGroup {
	if remove == 0 {
		return remaining[1:]
	}

	length := len(remaining)

	if remove == length-1 {
		return remaining[0:remove]
	}

	next := make([]*blokus.PieceGroup, length-1)

	copy(next[:], remaining[0:remove])
	copy(next[remove:], remaining[remove+1:])

	for i, n := range next {
		if n == nil {
			fmt.Printf("I HAVE FOUND NIL remove %v i %v\nremaining %v\nnext %v\n", remove, i, remaining, next)
		}
	}

	return next
}
