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

var seenMap = make(map[string]bool)
var lastSeen int

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	board := blokus.NewBoard(size)
	pieces := blokus.GetPieces()

	solve(board, pieces, make([]*usage, len(pieces)), 0, 0)
}

func solve(b *blokus.Board, remaining []*blokus.PieceGroup, used []*usage, level, count int) (bool, int) {

	if count-lastSeen > 500000 {
		lastSeen = count
		fmt.Printf("Reporting level %v & count %v\n%v\nseenMapLength: %v\n", level, count, b, len(seenMap))
	}

	done := false

	// Go through every remaining piece and attemp to use it in every possible
	// location.  As of now, I am not going to be smart about where / how I am
	// placing the pieces.  I am sure there is a way to be smarter, but I am
	// just not going to invest that time in yet.
	for i, pG := range remaining {

		// For every possible rotation
		for _, p := range pG.Pieces {

			// Attempt to place it on the board in every possible area
			maxRow := b.Row - p.Row
			maxCol := b.Col - p.Col

			for row := 0; row <= maxRow; row++ {
				for col := 0; col <= maxCol; col++ {

					// If I can add the piece to this location then i'll have
					// to recurse here and in the end remove the piece from this
					// location
					if b.Add(p, row, col) {
						used[level] = &usage{p, row, col}

						count++

						key := b.Key()

						// No need to continue down this path.
						if !b.IsSolvable() || seenMap[key] {
							b.Remove(p, row, col)
							continue
						}

						seenMap[key] = true

						// We have finished the function and the point of the
						// program!  Print the results and leave this forsaken
						// area
						if b.Solved() {
							fmt.Printf("We have done it in %v tries!\n", count)
							fmt.Printf("We have solved %vx%v!\n", b.Row, b.Col)

							fmt.Printf("Finished Board\n%v\n\n", b)
							for _, u := range used {
								fmt.Printf("Piece %v\n", u)
							}

							return true, count
						}

						// We are not solved then we need to recurse with the
						// next piece in the list
						clone := quickClone(remaining, i)

						// If solved, then return
						done, count = solve(b, clone, used, level+1, count)
						if done {
							return true, count
						}

						// Remove the piece we have currently placed and
						// continue on with the recursion
						b.Remove(p, row, col)
						used[level] = nil
					} // if add
				} // for col
			} // for row
		} // for piece
	} // for pieceGroup

	return false, count
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
