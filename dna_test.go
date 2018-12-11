package dna

import (
	"reflect"
	"testing"
)

func TestDNA_ReverseComplement(t *testing.T) {
	d := DNA{Sequence: []byte("ACCGG")}
	exp := DNA{Sequence: []byte("CCGGT")}
	got := d.ReverseComplement()
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("reverseComplemnt(%v)=%v; want %v", d, got, exp)
	}
}
