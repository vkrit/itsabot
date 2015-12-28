package goquery

import (
	"github.com/avabot/ava/Godeps/_workspace/src/github.com/andybalholm/cascadia"
	"github.com/avabot/ava/Godeps/_workspace/src/golang.org/x/net/html"
)

// Add adds the selector string's matching nodes to those in the current
// selection and returns a new Selection object.
// The selector string is run in the context of the document of the current
// Selection object.
func (s *Selection) Add(selector string) *Selection {
	return s.AddNodes(findWithMatcher([]*html.Node{s.document.rootNode}, cascadia.MustCompile(selector))...)
}

// AddMatcher adds the matcher's matching nodes to those in the current
// selection and returns a new Selection object.
// The matcher is run in the context of the document of the current
// Selection object.
func (s *Selection) AddMatcher(m Matcher) *Selection {
	return s.AddNodes(findWithMatcher([]*html.Node{s.document.rootNode}, m)...)
}

// AddSelection adds the specified Selection object's nodes to those in the
// current selection and returns a new Selection object.
func (s *Selection) AddSelection(sel *Selection) *Selection {
	if sel == nil {
		return s.AddNodes()
	}
	return s.AddNodes(sel.Nodes...)
}

// Union is an alias for AddSelection.
func (s *Selection) Union(sel *Selection) *Selection {
	return s.AddSelection(sel)
}

// AddNodes adds the specified nodes to those in the
// current selection and returns a new Selection object.
func (s *Selection) AddNodes(nodes ...*html.Node) *Selection {
	return pushStack(s, appendWithoutDuplicates(s.Nodes, nodes))
}

// AndSelf adds the previous set of elements on the stack to the current set.
// It returns a new Selection object containing the current Selection combined
// with the previous one.
func (s *Selection) AndSelf() *Selection {
	return s.AddSelection(s.prevSel)
}
