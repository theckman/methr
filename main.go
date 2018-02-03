package main

import (
	"fmt"

	"github.com/theckman/methr/methr"
)

type ifc interface {
	Get() int
	Set(i int)
}

func main() {
	np := methr.NonPointer{Value: 1}
	np.Set(42)
	fmt.Printf("np: %#v: %d\n", np, np.Get())

	p := &methr.Pointer{Value: 2}
	p.Set(42)
	fmt.Printf("n: %#v: %d\n", p, p.Get())

	b := methr.Both{Value: 3}
	b.Set(42)
	fmt.Printf("b: %#v: %d\n", b, b.Get())

	npp := methr.Pointer{Value: 4}
	npp.Set(42)
	fmt.Printf("npp: %#v: %d\n", npp, npp.Get())
}
