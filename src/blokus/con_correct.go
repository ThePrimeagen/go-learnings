package main

import (
	"fmt"

	"github.com/michaelbpaulson/blokus"
)

func solveCon(size int, used []*usage) []*usage {
	pieces := blokus.GetPieces()

	for i := range pieces {
		used := make([]*usage, len(pieces))
		p := rotate(pieces, i)

	}

	return nil
}

func rotate(p []*blokus.PieceGroup, point int) []*blokus.PieceGroup {
	newPieces := make([]*blokus.PieceGroup, len(p))

	count := len(p) - point
	newPieces = append(newPieces[:], p[:point]...)
	newPieces = append(newPieces[count:], p[point:]...)

	return newPieces
}

func _solveCon(b *blokus.Board, remaining []*blokus.PieceGroup, used []*usage, level, count int) (bool, int) {

	if !b.IsSolvable() {
		return false, count
	}

	if count%100000 == 0 {
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
						done, count = solveSC(b, clone, used, level+1, count)
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
