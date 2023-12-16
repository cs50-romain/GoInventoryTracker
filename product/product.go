package product

import "strings"

type Product struct {
	Name string
}

type Category struct {
	Name string
	Data []*Product
}

type Node struct {
	Category *Category
	left *Node
	right *Node
}

func Init() *Node {
	return &Node{&Category{}, nil, nil}
}

func (n *Node) Append(product, category string) {
	if n == nil {
		n = &Node{
			Category: &Category{category, nil},
			left: nil,
			right: nil,
		}

		n.Category.Data = append(n.Category.Data, &Product{product})
	} else {
		if strings.Compare(category, n.Category.Name) == -1 {
			n.left.Append(product, category)
		} else if strings.Compare(category, n.Category.Name) == 1 {
			// category is greater than n.category -> recurse to the right
			n.right.Append(product, category)
		} else if strings.Compare(category, n.Category.Name) == 0 {
			// Strings are equal, append product to Category.Data
			n.Category.Data = append(n.Category.Data, &Product{product})
		}
	}
}

func (n *Node) Remove(product, category string) {

}

func (n *Node) Search(product, category string) *Product {
	return &Product{}	
}
