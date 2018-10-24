package dna

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhred33ToQuality(t *testing.T) {
	assert.Equal(t, 2, Phred33ToQuality('#'))
	assert.Equal(t, 41, Phred33ToQuality('J'))
}
