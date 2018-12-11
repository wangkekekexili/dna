package dna

import "fmt"

// Read represents one read of DNA sequence.
type Read struct {
	Sequence []byte
	Quality  []byte
}

func (r *Read) String() string {
	return fmt.Sprintf("%s\n%s", string(r.Sequence), string(r.Quality))
}
