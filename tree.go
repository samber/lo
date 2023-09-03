package lo

import "fmt"

// TableToTree transforms database-like table into tree-structure
func TableToTree[R any, L any, K comparable](
	rows []R,
	rowToLeafFunc func(row *R, children []L) L,
	getKeysFunc func(row *R) (id K, parentId K, isRoot bool),
) ([]L, error) {

	type nodeType struct {
		payload  *R
		children []*nodeType
	}

	// Make node map
	nodeMap := make(map[K]*nodeType, len(rows))
	for _, row := range rows {
		row := row
		id, _, _ := getKeysFunc(&row)
		if _, exists := nodeMap[id]; exists {
			return nil, fmt.Errorf("lo.TableToTree: duplicate primary id")
		}
		nodeMap[id] = &nodeType{
			payload: &row,
		}
	}

	// Connect children nodes
	root := make([]*nodeType, 0)
	for _, row := range rows {
		id, parentId, isRoot := getKeysFunc(&row)
		node := nodeMap[id]

		if isRoot {
			root = append(root, node)
		} else {
			parentNode, exists := nodeMap[parentId]
			if !exists {
				return nil, fmt.Errorf("lo.TableToTree: bad parent id")
			}
			parentNode.children = append(parentNode.children, node)
		}
	}

	// Walk through the node tree and build leaf tree
	var walkFunc func([]*nodeType) []L
	walkFunc = func(nodes []*nodeType) []L {
		if len(nodes) == 0 {
			return nil
		}

		leafs := make([]L, 0, len(nodes))
		for _, node := range nodes {
			leaf := rowToLeafFunc(node.payload, walkFunc(node.children))
			leafs = append(leafs, leaf)
		}
		return leafs
	}

	return walkFunc(root), nil
}
