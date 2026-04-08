package tree_sitter_clojure_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_clojure "github.com/sogaiu/tree-sitter-clojure/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_clojure.Language())
	if language == nil {
		t.Errorf("Error loading Clojure grammar")
	}
}
