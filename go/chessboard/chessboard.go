package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

func CountInFile(cb Chessboard, file string) int {
	c := 0

	for _, v := range cb[file] {
		if v {
			c++
		}
	}

	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var c int
	if rank >= 1 && rank <= 8 {
		for _, v := range cb {
			if v[rank-1] {
				c++
			}
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var c int
	for _, v := range cb {
		for range v {
			c++
		}
	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var c int
	for _, v := range cb {
		for _, b := range v {
			if b {
				c++
			}
		}
	}
	return c
}
