package dna

type DNA struct {
	Sequence []byte
}

var complement = map[byte]byte{
	'A': 'T',
	'T': 'A',
	'G': 'C',
	'C': 'G',
	'N': 'N',
}

func (d DNA) ReverseComplement() DNA {
	result := make([]byte, len(d.Sequence))
	for i := range d.Sequence {
		result[len(d.Sequence)-1-i] = complement[d.Sequence[i]]
	}
	return DNA{Sequence: result}
}

func (d DNA) String() string {
	return string(d.Sequence)
}
