package octree

import (
	"fmt"
)

type Tree struct {
	Root Node
	Size uint64
}

func (t *Tree) String() string {
	return fmt.Sprintf("{Root:%v,Size:%d", t.Root, t.Size)
}

func (t *Tree) GetIndex(x, y, z int) ([]byte, uint64) {
	return nil, 0
}

func (t *Tree) Set(index []byte, length uint64, value interface{}) {
	var (
		s byte
		i uint64 = 0
		n *Ast
		m Node = t.Root
	)
	if uint64(len(index))*8/3 < length {
		panic("Possible index length smaller than selected length")
	}

	for ; i <= length; i++ {
		switch v := m.(type) {
		case *Ast:
			if i == length {
				n[s] = value
				return
			}
			// 1111 1111
			// Get 3 bits
			s = (index[i*3/8]<<byte(i%8))>>8 - byte(i%8)
			if v[s] == nil {
				v[s] = &Ast{}
			}
			n = v
			m = v[s]
		default:
			if i == length {
				n[s] = value
				return
			}
			l := &Ast{}
			n[s] = l
			n = l
			s = (index[i*3/8]<<byte(i%8))>>8 - byte(i%8)
			m = n[s]
		}
	}
}

func (t *Tree) Get(index []byte, length uint64) Node {
	var (
		i uint64 = 0
		n Node   = t.Root
	)
	if uint64(len(index))*8/3 < length {
		panic("Possible index length smaller than selected length")
	}
	for ; i <= length; i++ {
		switch v := n.(type) {
		case *Ast:
			s := (index[i*3/8]<<byte(i%8))>>8 - byte(i%8)
			n = v[s]
		default:
			return v
		}

	}
	return nil
}

type Node interface{}

type Ast [8]Node

func (a *Ast) String() string {
	return fmt.Sprintf("{0:%v,1:%v,2:%v,3:%v,4:%v,5:%v,6:%v,7:%v}", (*a)[0], (*a)[1], (*a)[2], (*a)[3], (*a)[4], (*a)[5], (*a)[6], (*a)[7])
}
