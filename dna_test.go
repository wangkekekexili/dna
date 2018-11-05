package dna

import (
	"reflect"
	"testing"
)

func TestDNA_ReverseComplement(t *testing.T) {
	d := DNA{Sequence: "ACCGG"}
	exp := DNA{Sequence: "CCGGT"}
	got := d.ReverseComplement()
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("reverseComplemnt(%v)=%v; want %v", d, got, exp)
	}
}
