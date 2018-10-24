package dna

func Phred33ToQuality(c int) int {
	return c - 33
}
