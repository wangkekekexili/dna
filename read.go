package dna

import "fmt"

// Read represents one read of DNA sequence.
type Read struct {
	Sequence string
	Quality  string
}

func (r *Read) String() string {
	return fmt.Sprintf("%s\n%s", r.Sequence, r.Quality)
}
