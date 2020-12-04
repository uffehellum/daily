package daily590

import (
	"fmt"
	"strings"
	"testing"
)

func TestGhost(t *testing.T) {
	actual := GhostFirstWinningPlays([]string{"ghost"})
	AssertEqualString(t,"",  actual, "ghost")
}

func TestCatCalfDogBear(t *testing.T) {
	actual := GhostFirstWinningPlays([]string{"cat", "calf", "dog", "bear"})
	AssertEqualString(t,"bc",  actual, "cat calf dog bear")
}

func AssertEqualString(t *testing.T, expected string, actual string, message string) {
	if actual != expected {
		t.Fatal("expected", expected, "actual", actual, message)
	}
}

type node struct {
	state int // -1 loss, 0, unknown, +1 win
	next map[rune] *node
}

func GhostFirstWinningPlays(dictionary []string) string {
	tree := createTree(dictionary)
	markWinnersAndLosers(tree)
	printNode(0, tree)
	return topLevelWinners(tree)
}

func topLevelWinners(tree *node) string {
	a := []rune{}
	for r, n := range tree.next {
		if n.state == 1 {
			a = append(a, r)
		}
	}
	return string(a)
}

func markWinnersAndLosers(tree *node) {
	if tree == nil {
		return
	}
	anyWinningChildren := false
	for _, n := range tree.next {
		markWinnersAndLosers(n)
		if n.state == 1 {
			anyWinningChildren = true
		}
	}
	if anyWinningChildren || len(tree.next) == 0 {
		tree.state = -1
	} else {
		tree.state = 1
	}
}

func createTree(dictionary []string) *node {
	tree := &node{next: make(map[rune]*node)}
	for _, word := range dictionary {
		insertWord(word, tree)
	}
	return tree
}

func printNode(tab int, n *node) {
	fmt.Println([]string{"lose", "unknown", "win"}[n.state + 1])
	for r, child := range n.next {
		fmt.Print(strings.Repeat("  ", tab), fmt.Sprintf("\"%c\" ", r))
		printNode(tab + 1, child)
	}
}

func insertWord(word string, tree *node) {
	p := tree
	for _, r := range word {
		_, ok := p.next[r]
		if !ok {
			p.next[r] = &node{next: make(map[rune]*node)}
		}
		p = p.next[r]
	}
}
