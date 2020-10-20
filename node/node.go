package simplesm

import "fmt"

type Node struct {
	ID       string
	children []*Node
}

func (p *Node) AddChild(child *Node) {
	if !contains(p.children, child.ID) {
		p.children = append(p.children, child)
	}
}

func (p *Node) draw(drawnNodes []*Node, globalDraw string) ([]*Node, string) {
	if contains(drawnNodes, p.ID) {
		return drawnNodes, globalDraw
	}

	drawnNodes = append(drawnNodes, p)
	for _, a := range p.children {
		globalDraw += fmt.Sprintf("\n %v --> %v", p.ID, a.ID)
	}

	for _, a := range p.children {
		drawnNodes, globalDraw = a.draw(drawnNodes, globalDraw)
	}
	return drawnNodes, globalDraw
}

func Draw(rootNode *Node) string {
	result := "@startuml"
	result += "\n[*] --> " + rootNode.ID
	_, result = rootNode.draw([]*Node{}, result)
	result += "\n@enduml"
	return result
}

func FindNode(rootNode *Node, ID string) *Node {
	return findNodeExcludeRecursivity([]*Node{}, rootNode, ID)
}

func IsValidTransition(rootNode *Node, from string, to string) bool {
	fromNode := FindNode(rootNode, from)
	if fromNode != nil {
		return contains(fromNode.children, to)
	}
	return false
}

func findNodeExcludeRecursivity(chequedNodes []*Node, node *Node, ID string) *Node {
	if node.ID == ID {
		return node
	}

	if contains(chequedNodes, node.ID) {
		return nil
	}

	chequedNodes = append(chequedNodes, node)

	for _, a := range node.children {
		if a.ID == ID {
			return a
		}
		b := findNodeExcludeRecursivity(chequedNodes, a, ID)
		if b != nil {
			return b
		}
	}
	return nil
}

func contains(arr []*Node, ID string) bool {
	for _, a := range arr {
		if a.ID == ID {
			return true
		}
	}
	return false
}
