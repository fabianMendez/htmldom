package htmldom

import "golang.org/x/net/html"

// GetAttribute returns the value for the specified attribute of the given node
func GetAttribute(n *html.Node, attrKey string) string {
	for _, a := range n.Attr {
		if a.Key == attrKey {
			return a.Val
		}
	}

	return ""
}

// GetElementMatching returns the firs node that matches the given predicate
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
