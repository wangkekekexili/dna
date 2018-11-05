package dna

type DNA struct {
	Sequence string
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
	for i := range []byte(d.Sequence) {
		result[len(d.Sequence)-1-i] = complement[d.Sequence[i]]
	}
	return DNA{Sequence: string(result)}
}
