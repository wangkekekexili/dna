package fastq

import (
	"bufio"
	"errors"
	"io"

	"github.com/wangkekekexili/dna"
)

var ErrNotFastq = errors.New("not a fastq file")

func Parse(r io.Reader) ([]*dna.Read, error) {
	var reads []*dna.Read

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var read dna.Read

		// Ignore first part.

		// Second part is sequence.
		if !scanner.Scan() {
			if scanner.Err() != nil {
				return nil, scanner.Err()
			} else {
				return nil, errors.New("")
			}
		}
		read.Sequence = scanner.Text()

		// Ignore third part.
		if !scanner.Scan() {
			if scanner.Err() != nil {
				return nil, scanner.Err()
			} else {
				return nil, errors.New("")
			}
		}

		// Forth part is quality.
		if !scanner.Scan() {
			if scanner.Err() != nil {
				return nil, scanner.Err()
			} else {
				return nil, errors.New("")
			}
		}
		read.Quality = scanner.Text()

		reads = append(reads, &read)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return reads, nil
}
