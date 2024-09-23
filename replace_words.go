package algorithms

import "strings"

//In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word derivative. For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".
//
//Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the derivatives in the sentence with the root forming it. If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.
//
//Return the sentence after the replacement.
//
//
//
//Example 1:
//
//Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
//Output: "the cat was rat by the bat"
//Example 2:
//
//Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
//Output: "a a b c"

func replaceWords(dictionary []string, sentence string) string {
	var sb strings.Builder
	trie := dictioryToTrie(dictionary)

	for _, word := range strings.Split(sentence, " ") {
		prefix := trie.MatchPrefix(word)

		if len(prefix) == 0 {
			sb.WriteString(" " + word)
		} else {
			sb.WriteString(" " + prefix)
		}
	}

	return strings.TrimSpace(sb.String())
}

func dictioryToTrie(dict []string) *Trie {
	trie := NewTrie()

	for i := range dict {
		trie.Insert(dict[i])
	}
	return trie
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.links[r]; !ok {
			newNode := &TrieNode{
				links:  make(map[rune]*TrieNode),
				isWord: false,
			}
			node.links[r] = newNode
		}
		node = node.links[r]
	}
	node.isWord = true
}

func (t *Trie) MatchPrefix(word string) string {
	var sb strings.Builder
	node := t.root
	for _, r := range word {
		if _, ok := node.links[r]; !ok {
			break
		}
		sb.WriteString(string(r))
		if node.links[r].isWord {
			return sb.String()
		}
		node = node.links[r]
	}
	if !node.isWord {
		return ""
	} else {
		return sb.String()
	}
}

func NewTrie() *Trie {
	node := &TrieNode{
		links: map[rune]*TrieNode{
			'-': nil,
		},
		isWord: false,
	}
	return &Trie{root: node}
}

type TrieNode struct {
	links  map[rune]*TrieNode
	isWord bool
}
