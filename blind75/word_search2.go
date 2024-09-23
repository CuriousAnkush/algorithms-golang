package blind75

func FindWords(board [][]string, words []string) []string {
	var answer []string
	for wIdx := range words {
		for row := range board {
			for col := range board[row] {
				if dfs2(board, row, col, words[wIdx], 0) {
					answer = append(answer, words[wIdx])
				}
			}
		}
		wIdx++
	}

	return answer
}

func dfs2(board [][]string, row int, column int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if string(word[index]) != board[row][column] {
		return false
	}
	oldVal := board[row][column]
	board[row][column] = "-"
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for idx := range directions {
		direction := directions[idx]

		if row+direction[0] < len(board) &&
			column+direction[1] < len(board[0]) &&
			row+direction[0] > -1 &&
			column+direction[1] > -1 {
			if dfs2(board, row+direction[0], column+direction[1], word, index+1) {
				return true
			}
		}
	}

	board[row][column] = oldVal

	return false
}

type TrieNode1 struct {
	isWordBreak bool
	links       map[rune]*TrieNode1
}

func createTrie() *TrieNode1 {
	return &TrieNode1{
		links: map[rune]*TrieNode1{
			'-': {
				isWordBreak: false,
				links:       make(map[rune]*TrieNode1),
			},
		},
		isWordBreak: false,
	}
}

func (n *TrieNode1) insertWord(word string) {
	root := n.links['-']
	for idx := range word {
		ch := rune(word[idx])
		if _, found := root.links[ch]; !found {
			node := &TrieNode1{
				links:       make(map[rune]*TrieNode1),
				isWordBreak: false,
			}
			root.links[ch] = node
		}
		root = root.links[ch]
	}
	root.isWordBreak = true
}

func (n *TrieNode1) findWords(prefix string) []string {
	var words []string
	root := n.links['-']
	for idx := range prefix {
		ch := rune(prefix[idx])
		if r, found := root.links[ch]; !found {
			return words
		} else {
			root = r
		}
	}
	recurFindWords(root.links, &words, prefix)

	return words
}

func recurFindWords(links map[rune]*TrieNode1, words *[]string, prefix string) {
	for key, tn := range links {
		word := prefix + string(key)
		if tn.isWordBreak {
			*words = append(*words, word)
		}
		recurFindWords(tn.links, words, word)
	}
}

func createDict(words []string) *TrieNode1 {
	tn := createTrie()
	for i := range words {
		tn.insertWord(words[i])
	}

	return tn
}
