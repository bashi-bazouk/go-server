package graph

import (
	"unsafe"
)

type Node int
type Edge struct {
	S Node
	T Node
}

type Object interface{}
type Arrow struct {
	D Object
	C Object
}

type Transformation struct {
	Mapfunc(Node, Object) (Node, Object)
	Reduce func(Edge, Arrow) (Edge, Arrow)
}


type AbstractGraph interface {
	GetArrow (Edge, *Object)
	SetArrow (Edge, *Object)
	DeleteArrow (Edge)
	MapReduce (Primitive, *AbstractGraph) error
}


func (g AbstractGraph) GetObject (object int) (*interface{}, bool) {
	return g.GetArrow(Arc { S: object, T: object })
}


func (g AbstractGraph) SetObject (object uint64, ident *interface{}) {
	g.SetArrow(object, object, ident)
}


func (g AbstractGraph) DeleteObject (object uint64) {
	delete(g, object)
	domains, _ := g.Neighbors(object)
	for _, domain := range domains {
		g.DeleteArrow(domain, object)
	}
}


func (g AbstractGraph) CoGraph () (cograph AbstractGraph) {
	for domain, codomains := range g {
		for codomain, ident := range codomains {
			cograph.SetArrow(codomain, domain, ident)
		}
	}
	return cograph
}


func (g AbstractGraph) Neighbors (object uint64) (domains []uint64, codomains []uint64) {
	for d, cs := range g {
		if d == object {
			codomains = cs
			continue
		} else {
			if _, has_codomain := cs[object]; has_codomain {
				append(domains, d)
			}
		}
	}
	return domains, codomains
}


func (g AbstractGraph) Slice (codomain uint64) (slice AbstractGraph) {
	domains, _ := g.Neighbors(codomain)
	for _, domain := range domains {
		slice.SetArrow(domain, codomain, g[domain][codomain])
	}
	return slice
}


func (g AbstractGraph) CoSlice (domain uint64) (coslice AbstractGraph) {
	_, codomains := g.Neighbors(domain)
	for _, codomain := range codomains {
		coslice.SetArrow(domain, codomain, g[domain][codomain])
	}
	return coslice
}


func (g AbstractGraph) Indices () (indices []uint64) {
	count := 0
	var reverseIndex map[uint64]int
	maybeAddIndex := func(index uint64) {
		if _, inMap := reverseIndex[index]; !inMap {
			reverseIndex[index] = count
			count = count + 1
		}
	}
	for domain, codomains := range g {
		maybeAddIndex(domain)
		for _, codomain := range codomains {
			maybeAddIndex(codomain)
		}
	}

	indices = indices[:count]
	for index, i := range reverseIndex {
		indices[i] = index
	}

	return indices
}


func FromObjects(objects []interface{}) (g AbstractGraph) {
	for _, object := range objects {
		pointer := *(*uint64)(unsafe.Pointer(&object))
		g.SetObject(pointer, object)
	}
	return g
}