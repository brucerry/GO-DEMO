package tree
import (
	"fmt"
	"sort"
)
// Record is the default struct for each entry
type Record struct {
	ID, Parent int
}
// Node is the default struct for parent nodes
type Node struct {
	ID       int
	Children []*Node
} 
// Build stitches together the node tree
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := map[int]*Node{}
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", r)
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}