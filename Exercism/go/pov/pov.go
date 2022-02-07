package pov

// Define the Graph type here.
type Graph struct {
    nodes []*Node
}

type Node struct {
    label string
    from string
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(nodeLabel string) {
    node := Node{nodeLabel, ""}
    for _,n := range g.nodes {
        if n.label == nodeLabel {
            return
        }
    }
    g.nodes = append(g.nodes, &node)
}

func (g *Graph) AddArc(from, to string) {
    addNode := true
	for _,n := range g.nodes {
        if n.label == from {
            addNode = false
            break
        }
    }
    if addNode {
        g.AddNode(from)
    }
	for _,n := range g.nodes {
        if n.label == to {
            n.from = from
            return
        }
    }
}

func (g *Graph) ArcList() []string {
    s := []string{}
	for _,n := range g.nodes {
        if n.from != "" && n.label != "" {
            s = append(s, n.from + " -> " + n.label)
        }
    }
    return s
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
    new_g := New()
    target := newRoot
	for {
        matched := false // "matched" means g still has a node that has to be re-directed
        for i,n := range g.nodes {
            if n.label == target {
                new_g.AddNode(target)
                new_g.AddNode(n.from)
                new_g.AddArc(target, n.from) // add node & arc into new_g in new direction
                target = n.from
                g.nodes = append(g.nodes[:i], g.nodes[i+1:]...) // delete the node from g
                matched = true
                break
            }
        }
        if !matched {
            break
        }
    }

    // process the nodes that remained in g
    // i.e. add all of them into new_g
    for _,n := range g.nodes {
        new_g.AddNode(n.label)
        new_g.AddArc(n.from, n.label)
    }

    return new_g
}
