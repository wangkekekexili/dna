package fasta

import (
	"bufio"
	"errors"
	"io"

	"github.com/wangkekekexili/dna"
)

var ErrNotFasta = errors.New("not a fasta file")

func Parse(r io.Reader) (*dna.DNA, error) {
	d := &dna.DNA{}
	scanner := bufio.NewScanner(r)

	// Ignore first line.
	if !scanner.Scan() {
		return nil, ErrNotFasta
	}
	for scanner.Scan() {
		d.Sequence += scanner.Text()
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return d, nil
}
