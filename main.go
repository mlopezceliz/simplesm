package main

import (
	"fmt"

	simplesm "github.com/mlopezceliz/simplesm/node"
)

func main() {
	a := &simplesm.Node{ID: "a"}
	b := &simplesm.Node{ID: "b"}
	c := &simplesm.Node{ID: "c"}
	d := &simplesm.Node{ID: "d"}
	e := &simplesm.Node{ID: "e"}

	a.AddChild(b)
	a.AddChild(c)
	b.AddChild(d)
	d.AddChild(e)
	e.AddChild(a)

	isValid := simplesm.IsValidTransition(a, "b", "d")
	fmt.Printf("\n Es valido el cambio de estado de b a d -> %v", isValid)

	isValid = simplesm.IsValidTransition(a, "d", "b")
	fmt.Printf("\n Es valido el cambio de estado de d a b -> %v", isValid)

	isValid = simplesm.IsValidTransition(a, "a", "e")
	fmt.Printf("\n Es valido el cambio de estado de a a e -> %v", isValid)

	isValid = simplesm.IsValidTransition(a, "e", "a")
	fmt.Printf("\n Es valido el cambio de estado de e a a -> %v", isValid)

	diagram := simplesm.Diagram(a)
	fmt.Printf("\n Diagrama -> %v", diagram)
	link := simplesm.Link(a)
	fmt.Printf("\n Link -> %v", link)

}
