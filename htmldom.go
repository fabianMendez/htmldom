package htmldom

import (
	"strings"

	"golang.org/x/net/html"
)

// GetAttribute returns the value for the specified attribute of the given node
func GetAttribute(n *html.Node, attrKey string) string {
	for _, a := range n.Attr {
		if a.Key == attrKey {
			return a.Val
		}
	}

	return ""
}

// GetElementMatching returns the first node that matches the given predicate
func GetElementMatching(node *html.Node, fn func(*html.Node) bool) *html.Node {
	if fn(node) {
		return node
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if elm := GetElementMatching(child, fn); elm != nil {
			return elm
		}
	}

	return nil

}

func GetAllElementsMatching(node *html.Node, fn func(*html.Node) bool) []*html.Node {
	var elms []*html.Node

	if fn(node) {
		elms = append(elms, node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		elms = append(elms, GetAllElementsMatching(child, fn)...)
	}

	return elms
}

// GetElementByID returns the node with the specified id or nil if it can not be found
func GetElementByID(node *html.Node, id string) *html.Node {
	return GetElementMatching(node, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			elmID := GetAttribute(n, "id")
			return elmID == id
		}

		return false
	})
}

// ContainsClass returns whether the node contains the specified class
func ContainsClass(n *html.Node, class string) bool {
	if n.Type == html.ElementNode {
		elmClassAttr := GetAttribute(n, "class")
		elmClasses := strings.Split(elmClassAttr, " ")
		for _, elmClass := range elmClasses {
			if elmClass == class {
				return true
			}
		}
	}

	return false
}

// GetAllElementsByClass returns every node that has the specified class
func GetAllElementsByClass(node *html.Node, class string) []*html.Node {
	return GetAllElementsMatching(node, func(n *html.Node) bool {
		return ContainsClass(n, class)
	})
}

// GetElementByClass returns the first node containing the specified class
func GetElementByClass(node *html.Node, class string) *html.Node {
	return GetElementMatching(node, func(n *html.Node) bool {
		return ContainsClass(n, class)
	})
}

// GetInnerText returns the inner text of the node
func GetInnerText(n *html.Node) string {
	str := ""

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			str += c.Data
		} else if IsTag(c, "br") {
			str += "\n"
		} else if !IsTag(c, "script") {
			str += GetInnerText(c)
		}
	}

	return str
}

// IsTag returns true if the node is of the given tag
func IsTag(node *html.Node, tag string) bool {
	return node.Type == html.ElementNode && node.Data == tag
}

// GetAllElementsByTag returns all the nodes of the given tag
func GetAllElementsByTag(node *html.Node, tag string) []*html.Node {
	return GetAllElementsMatching(node, func(n *html.Node) bool {
		return IsTag(n, tag)
	})
}

// GetElementByTag returns the first element of the given tag
func GetElementByTag(node *html.Node, tag string) *html.Node {
	return GetElementMatching(node, func(n *html.Node) bool {
		return IsTag(n, tag)
	})
}
