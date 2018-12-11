package proximate

import "testing"

func TestEditDistance(t *testing.T) {
	tests := []struct {
		left, right []byte
		expDistance int
	}{
		{nil, nil, 0},
		{[]byte{}, []byte{'A', 'C'}, 2},
		{[]byte{'A'}, []byte{'C'}, 1},
		{[]byte{'A'}, []byte{'C', 'A'}, 1},
		{[]byte{'A', 'G'}, []byte{'C', 'T'}, 2},
		{[]byte{'A', 'C', 'G', 'T'}, []byte{'A', 'T'}, 2},
	}
	for _, test := range tests {
		gotDistance := EditDistance(test.left, test.right)
		if gotDistance != test.expDistance {
			t.Fatalf("got distance %v; want %v", gotDistance, test.expDistance)
		}
	}
}
