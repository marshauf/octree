package octree

import (
	"testing"
)

func TestByteIndexed(t *testing.T) {
	tree := &Tree{Root: &Ast{
		0: &Leaf{},
		1: &Leaf{},
	}}
	v := tree.Get([]byte{0}, 1)
	if v == nil {
		t.Error("Did not find node 0")
	}
	v = tree.Get([]byte{1}, 1)
	if v == nil {
		t.Error("Did not find node 1")
	}
}

func TestByteIndexedRandom(t *testing.T) {
	indexes := [][]byte{
		[]byte{0},
		[]byte{1},
		[]byte{2},
		[]byte{3},
		[]byte{255},
	}

	for i := range indexes {
		tree := &Tree{Root: &Ast{}}
		var l uint64 = 1
		tree.Set(indexes[i], l, "test")
		v := tree.Get(indexes[i], l)
		if v != "test" {
			t.Logf("l:%d", l)
			t.Errorf("%v", tree)
		}
	}
}

type Leaf struct {
}
